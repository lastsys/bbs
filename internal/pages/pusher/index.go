package pusher

import (
	"errors"
	"fmt"
	"github.com/lastsys/bbs/internal/screen"
	"log"
	"strconv"
	"time"

	"github.com/lastsys/bbs/internal/pages"
	"github.com/lastsys/bbs/internal/pages/util"
	"github.com/lastsys/bbs/internal/protocol"
	"github.com/lastsys/bbs/internal/user"
	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/parser"
)

func Index(s *user.Session) {
	board := initializeBoard()
	robots := initializeRobots()

	s.Buffer.Clear()
	util.WriteCombineName(s)
	util.WriteCombineLogo(s)
	//for i := 0; i < 10; i++ {
	buildBoard(board, robots)
	drawBoard(s, board)
	writeStatus(s, robots)
	s.UpdateClient()

	//redActions := runScript("assets/default.js", BlueTeam, board, robots)
	//blueActions := runScript("assets/default.js", RedTeam, board, robots)
	//
	//robots.applyActions(redActions, board)
	//robots.applyActions(blueActions, board)

	//time.Sleep(500 * time.Millisecond)
	//}

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

func drawBoard(s *user.Session, board *Board) {
	for row, tiles := range *board {
		for col, tile := range tiles {
			switch tile {
			case Floor:
				s.Buffer.Write(floorTile, uint8(row)+1, uint8(col)+1)
			case Wall:
				s.Buffer.Write(wallTile, uint8(row)+1, uint8(col)+1)
			case RedRobot1, RedRobot2, RedRobot3, RedRobot4, RedRobot5:
				s.Buffer.Write(redRobotTile, uint8(row)+1, uint8(col)+1)
			case BlueRobot1, BlueRobot2, BlueRobot3, BlueRobot4, BlueRobot5:
				s.Buffer.Write(blueRobotTile, uint8(row)+1, uint8(col)+1)
			}
		}
	}
}

var halt = errors.New("script timeout")

const ScriptTimeout = 5

// Run robot script in a goroutine with timeout.
func runScript(path string, team Team, board *Board, robots *Robots) *Actions {
	actionChannel := make(chan Actions)

	vm := otto.New()
	vm.Interrupt = make(chan func(), 1) // Non-blocking with buffer.
	injectData(vm, robots, board)
	program, err := parser.ParseFile(nil, path, nil, 0)
	if err != nil {
		log.Println(err)
		return nil
	}

	go func(channel chan Actions, vm *otto.Otto) {
		vm.Run(program)
		result, err := vm.Call(`run`, nil)
		if err != nil {
			log.Println(err)
			channel <- nil
			return
		}
		channel <- jsActionsToGoActions(result)
	}(actionChannel, vm)

	select {
	case result := <-actionChannel:
		return &result
	case <-time.After(ScriptTimeout * time.Second):
		log.Printf("Timeout after %v seconds.\n", ScriptTimeout)
		vm.Interrupt <- func() {
			panic(halt)
		}
		return nil
	}
}

func injectData(vm *otto.Otto, robots *Robots, board *Board) {
	if value, err := vm.ToValue(robots); err == nil {
		if err := vm.Set("ROBOTS", value); err != nil {
			log.Println("Failed to set ROBOTS:", err)
		}
	} else {
		log.Println("Failed to convert robots to value.")
	}
	if value, err := vm.ToValue(board); err == nil {
		if err := vm.Set("BOARD", value); err != nil {
			log.Println("Failed to set BOARD:", err)
		}
	} else {
		log.Println("Failed to convert board to value.")
	}
	if err := vm.Set("BOARD_WIDTH", BoardWidth); err != nil {
		log.Println("Failed to set BOARD_WIDTH.")
	}
	if err := vm.Set("BOARD_HEIGHT", BoardHeight); err != nil {
		log.Println("Failed to set BOARD_HEIGHT.")
	}
}

// Convert JS-structure to list of Action.
func jsActionsToGoActions(value otto.Value) Actions {

	array := value.Object()

	actionArray := make(Actions)

	for _, i := range array.Keys() {
		item, err := array.Get(i)
		if err != nil {
			log.Println("Failed to get item", i)
			return nil
		}
		actionObject := item.Object()
		dxValue, err := actionObject.Get("dx")
		if err != nil {
			log.Println("Failed to get dx.")
			return nil
		}
		dyValue, err := actionObject.Get("dy")
		if err != nil {
			log.Println("Failed to get dy.")
			return nil
		}

		robotId, err := strconv.Atoi(i)
		if err != nil {
			log.Println("Failed to convert robotId to integer.")
			return nil
		}
		dx, err := dxValue.ToInteger()
		if err != nil {
			log.Println("Failed to convert dx to integer.")
			return nil
		}
		dy, err := dyValue.ToInteger()
		if err != nil {
			log.Println("Failed to convert dy to integer.")
			return nil
		}

		actionArray[RobotId(robotId)] = &Action{uint8(dx), uint8(dy)}
	}

	return actionArray
}

func writeStatus(s *user.Session, robots *Robots) {
	blueRobotsAlive := 0
	redRobotsAlive := 0
	for id := BlueRobot1; id <= BlueRobot5; id++ {
		robot, ok := (*robots)[RobotId(id)]
		if ok && !robot.dead {
			blueRobotsAlive++
		}
	}
	for id := RedRobot1; id <= RedRobot5; id++ {
		robot, ok := (*robots)[RobotId(id)]
		if ok && !robot.dead {
			redRobotsAlive++
		}
	}
	s.Buffer.Print(fmt.Sprintf("Blue Team : %v", blueRobotsAlive), 1, 23, screen.LightBlue, screen.Black)
	s.Buffer.Print(fmt.Sprintf("Red Team  : %v", redRobotsAlive), 2, 23, screen.LightRed, screen.Black)
}
