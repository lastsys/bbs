package pusher

import "github.com/lastsys/bbs/internal/screen"

const (
	BoardWidth  = 21
	BoardHeight = 21
)

type Tile uint8

type Board [][]Tile

const (
	Floor      Tile = 0x00
	Wall       Tile = 0x01
	RedRobot1  Tile = 0x10
	RedRobot2  Tile = 0x11
	RedRobot3  Tile = 0x12
	RedRobot4  Tile = 0x13
	RedRobot5  Tile = 0x14
	BlueRobot1 Tile = 0x80
	BlueRobot2 Tile = 0x81
	BlueRobot3 Tile = 0x82
	BlueRobot4 Tile = 0x83
	BlueRobot5 Tile = 0x84
)

var level = []string{
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
var redRobotTile = screen.Character{81, screen.LightRed, screen.DarkGray, false}
var blueRobotTile = screen.Character{81, screen.LightBlue, screen.DarkGray, false}

func initializeBoard() *Board {
	rows := make([][]Tile, BoardHeight)
	for i := range rows {
		rows[i] = make([]Tile, BoardWidth)
	}
	board := Board(rows)
	return &board
}

func buildBoard(board *Board, robots *Robots) {
	// Basic level information.
	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			switch level[row][col] {
			case 'x':
				(*board)[row][col] = Wall
			case '.':
				(*board)[row][col] = Floor
			}
		}
	}
	// Place robots.
	for _, robot := range *robots {
		(*board)[robot.position.row][robot.position.column] = Tile(robot.id)
	}
}
