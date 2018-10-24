package util

import (
	"github.com/lastsys/bbs/internal/screen"
	"github.com/lastsys/bbs/internal/user"
)

var blocks = [][]uint8{
	{22, 34},
	{23, 34},
	{19, 36},
	{20, 36},
	{21, 36},
	{22, 36},
	{23, 36},
	{18, 38},
	{19, 38},
	{21, 38},
	{23, 38},
}

var diags = [][]uint8{
	{21, 34},
	{17, 38},
}

func WriteCombineLogo(s *user.Session) {
	var (
		block = screen.Character{32, screen.Black, screen.White, false}
		diag  = screen.Character{105, screen.Black, screen.White, false}
	)

	for _, coord := range blocks {
		s.Buffer.Write(block, coord[0], coord[1])
	}

	for _, coord := range diags {
		s.Buffer.Write(diag, coord[0], coord[1])
	}
}

func WriteCombineName(s *user.Session) {
	s.Buffer.Print("COMBINE", 23, 26, screen.White, screen.Black)
}
