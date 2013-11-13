package main

import (
	"github.com/mriehl/go-tdd-katas/go/rover"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "[RoverDemo] ", log.Flags())
	r := rover.New(rover.Coordinates{99, 99}, rover.North)
	r.Grid.Insert(rover.Coordinates{99, 1}, rover.OBSTACLE)

	remote := rover.NewRemoteControl(r)
	remote.Send("F", "F")
	logger.Println("Rover is now at ", r.Coords)
}
