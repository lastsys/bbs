package welcome

import (
	"fmt"
	"github.com/lastsys/bbs/internal/pages"
	"github.com/lastsys/bbs/internal/pages/util"
	"github.com/lastsys/bbs/internal/protocol"
	"github.com/lastsys/bbs/internal/screen"
	"github.com/lastsys/bbs/internal/user"
)

var menuItems = []string{
	"Pusher",
	"Game of Life",
	"About Combine",
}

func Index(s *user.Session) {
	s.Buffer.Clear()
	util.WriteCombineName(s)
	util.WriteCombineLogo(s)

	var selectedItem = 0
	writeMenu(s, selectedItem)

	s.UpdateClient()

OuterLoop:
	for {
		msg := <-s.MessageChannel
		switch msg.(type) {
		case protocol.KeyCode:
			if keyCode, ok := msg.(protocol.KeyCode); ok {
				switch keyCode {
				case protocol.UpArrow:
					selectedItem = selectedItem - 1
					if selectedItem < 0 {
						selectedItem = len(menuItems) - 1
					}
					writeMenu(s, selectedItem)
					s.UpdateClient()
				case protocol.DownArrow:
					selectedItem = (selectedItem + 1) % len(menuItems)
					writeMenu(s, selectedItem)
					s.UpdateClient()
				case protocol.Enter:
					break OuterLoop
				}
			}
		}
	}

	switch selectedItem {
	case 0:
		s.Navigate(pages.Pusher)
	case 1:
		s.Navigate(pages.Life)
	case 2:
		s.Navigate(pages.About)
	}
}

func writeMenu(s *user.Session, selectedItem int) {
	var space = screen.Character{32, screen.White, screen.DarkGray, false}
	var row, col uint8
	const boxWidth = 30
	const menuStart = 4
	const menuCol = 3

	for row = 1; row <= menuStart+uint8(len(menuItems)); row++ {
		for col = 1; col <= boxWidth; col++ {
			s.Buffer.Write(space, row, col)
		}
	}

	s.Buffer.Print("Please make a choice:", 2, 2, screen.White, screen.DarkGray)

	color := screen.DarkGray
	for i, str := range menuItems {
		if selectedItem == i {
			color = screen.Gray
		} else {
			color = screen.DarkGray
		}
		s.Buffer.Print(fmt.Sprintf("%v. %v", i+1, str), menuStart+uint8(i), menuCol,
			screen.White, color)
	}
}
