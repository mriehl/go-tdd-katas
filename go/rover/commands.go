package rover

import (
	"fmt"
	"log"
	"os"
)

var rcLogger = log.New(os.Stderr, "[RoverRemoteControl] ", log.Flags())

func ValidateCommands(commands []string) bool {
	for _, command := range commands {
		switch command {
		case "F":
		case "B":
		case "L":
		case "R":
		default:
			rcLogger.Println(fmt.Sprintf("Unknown command: %q, ignoring it.", command))
			return false
		}
	}
	return true
}

type RemoteControl struct {
	Rover *Rover
}

func (remote *RemoteControl) Send(commands ...string) bool {
	if !ValidateCommands(commands) {
		return false
	}

	for _, command := range commands {
		switch command {
		case "F":
			remote.Rover.Advance()
		case "B":
			remote.Rover.Retreat()
		case "L":
			remote.Rover.TurnLeft()
		case "R":
			remote.Rover.TurnRight()
		}
	}

	return true
}

func NewRemoteControl(rover *Rover) *RemoteControl {
	remote := new(RemoteControl)
	remote.Rover = rover
	return remote
}
