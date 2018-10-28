package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/lastsys/bbs/internal/pages/welcome"
	"github.com/lastsys/bbs/internal/user"
	"golang.org/x/net/http2"
)

const (
	staticPath            = "/static/"
	socketReadBufferSize  = 1024
	socketWriteBufferSize = 1024
	writeTimeout          = 15
	readTimeout           = 15
)

type socketRegister struct {
	set   map[*websocket.Conn]bool
	mutex sync.Mutex
}

var sockets socketRegister

// Store a websocket connection in the register.
func (s *socketRegister) register(conn *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.set[conn] = true
	log.Println("Current connections:", s.Count())
}

// Remove a websocket connection from the register.
func (s *socketRegister) unregister(conn *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.set, conn)
	log.Println("Current connections:", s.Count())
}

// Number of websockets in register.
func (s *socketRegister) Count() int {
	return len(s.set)
}

// Start server and initialize all handlers.
func StartServer(url string) {
	sockets.set = make(map[*websocket.Conn]bool)

	router := initRouter()
	server := http.Server{
		Handler:      router,
		Addr:         url,
		WriteTimeout: writeTimeout * time.Second,
		ReadTimeout:  readTimeout * time.Second,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	log.Println("Starting server at", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/ws", webSocketHandler)
	r.PathPrefix(staticPath).
		Handler(http.StripPrefix(staticPath, http.FileServer(http.Dir("./web/static/"))))
	return r
}

// Handle root path and return index.html.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	index, err := ioutil.ReadFile("./web/static/index.html")
	if err != nil {
		log.Println("Failed to read index.html")
		return
	}
	w.Write(index)
}

// Upgrade request to websocket.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  socketReadBufferSize,
	WriteBufferSize: socketWriteBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handle websocket request and upgrade connection.
func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	sockets.register(conn)

	// Make session.
	clientSession := user.NewSession(conn)
	go messageHandler(conn, clientSession.MessageChannel)
	go navigationHandler(clientSession)
	go welcome.Index(clientSession)
}
