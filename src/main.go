package main

import (
	"ObjectOrientedCourse/src/Player"
	"ObjectOrientedCourse/src/Rooms"
	"fmt"
)

func RunCommandLoop() {
	hasMoved := false
	for !hasMoved {
		availableCommandString := Player.GetAvailableCommandString()
		fmt.Println("Enter command (" + availableCommandString + "):")
		var command string
		fmt.Scanln(&command)
		Player.RunCommand(command, &hasMoved)
	}
}

func main() {
	Rooms.Initialize()
	Player.Initialize()
	for {
		Player.DescribeRoom()
		RunCommandLoop()
		if !Player.Alive() {
			break
		}
	}
}
