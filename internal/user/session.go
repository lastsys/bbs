package user

import (
	"github.com/gorilla/websocket"
	"github.com/lastsys/bbs/internal/pages"
	"github.com/lastsys/bbs/internal/screen"
)

const bufferSize = 16

type Session struct {
	Socket            *websocket.Conn
	Buffer            screen.Buffer
	MessageChannel    chan interface{}
	NavigationChannel chan pages.Page
}

func (s *Session) UpdateClient() {
	// Send update.
	// TODO: Optimize by sending differences only.
	s.Socket.WriteMessage(websocket.BinaryMessage, s.Buffer.SerializeMessage())
}

func (s *Session) Navigate(p pages.Page) {
	s.NavigationChannel <- p
}

func NewSession(socket *websocket.Conn) *Session {
	session := &Session{}
	session.Socket = socket
	session.Buffer.Clear()
	session.MessageChannel = make(chan interface{}, bufferSize)
	session.NavigationChannel = make(chan pages.Page)
	return session
}
