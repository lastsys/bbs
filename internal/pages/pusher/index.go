package pusher

import (
	"github.com/lastsys/bbs/internal/pages"
	"github.com/lastsys/bbs/internal/pages/util"
	"github.com/lastsys/bbs/internal/protocol"
	"github.com/lastsys/bbs/internal/screen"
	"github.com/lastsys/bbs/internal/user"
)

func Index(s *user.Session) {
	s.Buffer.Clear()
	util.WriteCombineName(s)
	util.WriteCombineLogo(s)
	s.Buffer.Print("Not Implemented...", 1, 1, screen.White, screen.Black)
	s.UpdateClient()

OuterLoop:
	for {
		msg := <-s.MessageChannel
		switch msg.(type) {
		case protocol.KeyCode:
			if keyCode, ok := msg.(protocol.KeyCode); ok {
				switch keyCode {
				case protocol.Enter:
					break OuterLoop
				}
			}
		}
	}

	s.Navigate(pages.Welcome)
}
