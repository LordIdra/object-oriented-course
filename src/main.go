package main

import (
	"ObjectOrientedCourse/src/Rooms"
)

var gameRunning bool
var room *Rooms.Room

func main() {
	Rooms.Initialize()
	room = Rooms.DefaultRoom()
	for gameRunning {

	}
}
