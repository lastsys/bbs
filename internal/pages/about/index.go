package about

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
	s.Buffer.Print("We are a consulting company built by", 1, 1, screen.White, screen.Black)
	s.Buffer.Print("social nerds with a passion for", 2, 1, screen.White, screen.Black)
	s.Buffer.Print("helping our customers take their", 3, 1, screen.White, screen.Black)
	s.Buffer.Print("business to the next level.", 4, 1, screen.White, screen.Black)

	s.Buffer.Print("At Combine, we are experts in control", 6, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("systems, model-based design and data", 7, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("science. Since the company was founded", 8, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("in 2002 we have been supplying", 9, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("professional engineering services and", 10, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("solutions, developed in the field or", 11, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("in-house. We also provide specialist", 12, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("or supplementary customer training.", 13, 1, screen.LightGray, screen.Black)

	s.Buffer.Print("Our vision is to improve technology", 15, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("around the world. Or as we like to", 16, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("say: \"Enter the next level\". We", 17, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("follow this motto by helping our", 18, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("customers' success, and we make", 19, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("progress by combining our", 20, 1, screen.LightGray, screen.Black)
	s.Buffer.Print("strengths.", 21, 1, screen.LightGray, screen.Black)

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
