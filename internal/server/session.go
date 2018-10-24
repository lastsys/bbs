package server

import (
	"github.com/gorilla/websocket"
	"github.com/lastsys/bbs/internal/pages"
	"github.com/lastsys/bbs/internal/pages/about"
	"github.com/lastsys/bbs/internal/pages/pusher"
	"github.com/lastsys/bbs/internal/pages/welcome"
	"github.com/lastsys/bbs/internal/protocol"
	"github.com/lastsys/bbs/internal/user"
	"log"
)

// Read messages coming from clients through WebSockets, parse and send the result to the given channel.
func messageHandler(s *websocket.Conn, c chan interface{}) {
	for {
		messageType, msg, err := s.ReadMessage()
		if err != nil {
			log.Println(err)
			log.Println("Unregistering socket.")
			sockets.unregister(s)
			close(c)
			log.Println(err)
			return
		}

		switch messageType {
		case websocket.BinaryMessage:
			if parsedMsg := protocol.ParseMessage(msg); parsedMsg != nil {
				c <- parsedMsg
			}
		case websocket.TextMessage:
			log.Println("Received unexpected Text Message.")
		}
	}
}

func navigationHandler(clientSession *user.Session) {
	for page := range clientSession.NavigationChannel {
		log.Println("Go to page", page)
		switch page {
		case pages.Welcome:
			go welcome.Index(clientSession)
		case pages.Pusher:
			go pusher.Index(clientSession)
		case pages.About:
			go about.Index(clientSession)
		}
	}
}
