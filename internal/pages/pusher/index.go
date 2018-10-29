package pusher

import (
	"github.com/lastsys/bbs/internal/pages"
	"github.com/lastsys/bbs/internal/pages/util"
	"github.com/lastsys/bbs/internal/protocol"
	"github.com/lastsys/bbs/internal/screen"
	"github.com/lastsys/bbs/internal/user"
)

const (
	BoardWidth  = 21
	BoardHeight = 21
)

var board = []string{
	".........xxx.........",
	".....................",
	"..xx.............xx..",
	"..x...............x..",
	".....................",
	".....................",
	".....................",
	".....................",
	"..........x..........",
	"x.........x.........x",
	"x.......xxxxx.......x",
	"x.........x.........x",
	"..........x..........",
	".....................",
	".....................",
	".....................",
	".....................",
	"..x...............x..",
	"..xx.............xx..",
	".....................",
	".........xxx.........",
}

var floorTile = screen.Character{32, screen.White, screen.DarkGray, false}
var wallTile = screen.Character{32, screen.Black, screen.LightGray, false}

func Index(s *user.Session) {
	s.Buffer.Clear()
	util.WriteCombineName(s)
	util.WriteCombineLogo(s)
	drawEmptyBoard(s)
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

func drawEmptyBoard(s *user.Session) {
	for row, tiles := range board {
		for col, tile := range tiles {
			switch tile {
			case '.':
				s.Buffer.Write(floorTile, uint8(row)+1, uint8(col)+1)
			case 'x':
				s.Buffer.Write(wallTile, uint8(row)+1, uint8(col)+1)
			}
		}
	}
}

type Position struct {
	row    uint8
	column uint8
}

type Role uint8

const (
	HunterAgent Role = iota
	PreyAgent
)

type Agent struct {
	position Position
	role     Role
}

type Agents []Agent

var agents = Agents{
	Agent{Position{1, 1}, HunterAgent},
	Agent{Position{20, 20}, PreyAgent},
}

// Move agents according to rules.
func (a *Agents) Update(board []string) {

}
