package pusher

import "log"

type Position struct {
	row    uint8
	column uint8
}

type Team uint8

const (
	BlueTeam Team = 0x00
	RedTeam  Team = 0x01
)

type RobotId uint8

type Robot struct {
	position Position
	id       RobotId
	team     Team
	dead     bool
}

func (r *Robot) applyAction(a *Action, board *Board) {
	if r.dead {
		return
	}

	newColumn := r.position.column + a.dx
	newRow := r.position.row + a.dy

	if newColumn >= BoardWidth || newRow >= BoardHeight {
		r.dead = true
		return
	}

	if (*board)[newRow][newColumn] == Floor {
		r.position.column = newColumn
		r.position.row = newRow
	}
}

type Robots map[RobotId]*Robot
type Actions map[RobotId]*Action

func (r *Robots) applyActions(actions *Actions, board *Board) {
	for robotId, action := range *actions {
		robot, ok := (*r)[robotId]
		if !ok {
			log.Println("Could not get robot", robotId)
		}
		robot.applyAction(action, board)
	}
}

func initializeRobots() *Robots {
	robots := Robots{
		RobotId(RedRobot1):  &Robot{Position{1, 1}, RobotId(RedRobot1), RedTeam, false},
		RobotId(RedRobot2):  &Robot{Position{1, 2}, RobotId(RedRobot2), RedTeam, false},
		RobotId(RedRobot3):  &Robot{Position{1, 3}, RobotId(RedRobot3), RedTeam, false},
		RobotId(RedRobot4):  &Robot{Position{1, 4}, RobotId(RedRobot4), RedTeam, false},
		RobotId(RedRobot5):  &Robot{Position{1, 5}, RobotId(RedRobot5), RedTeam, false},
		RobotId(BlueRobot1): &Robot{Position{19, 19}, RobotId(BlueRobot1), BlueTeam, false},
		RobotId(BlueRobot2): &Robot{Position{19, 18}, RobotId(BlueRobot2), BlueTeam, false},
		RobotId(BlueRobot3): &Robot{Position{19, 17}, RobotId(BlueRobot3), BlueTeam, false},
		RobotId(BlueRobot4): &Robot{Position{19, 16}, RobotId(BlueRobot4), BlueTeam, false},
		RobotId(BlueRobot5): &Robot{Position{19, 15}, RobotId(BlueRobot5), BlueTeam, false},
	}
	return &robots
}

type Action struct {
	dx uint8
	dy uint8
}
