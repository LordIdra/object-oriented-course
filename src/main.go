package main

import (
	"fmt"
	"ObjectOrientedCourse/src/Rooms"
	"ObjectOrientedCourse/src/Commands"
)

var gameRunning bool
var room *Rooms.Room

func RunCommandLoop() {
	roomChanged := false
	for !roomChanged {
		fmt.Println("Enter command (talk, fight, move):")
		var command string
		fmt.Scanln(&command)
		if command == "talk" {
			Commands.CommandTalk(room)
		} else if command == "fight" {
			if !Commands.CommandFight(room) {
				gameRunning = false
			}
		} else if command == "move" {
			room = Commands.CommandMove(room)
			roomChanged = true
		} else {
			fmt.Println("Invalid command!")
		}
	}
}

func main() {
	Rooms.Initialize()
	room = Rooms.DefaultRoom()
	for gameRunning {
		room.Describe()
		RunCommandLoop()
	}
}
