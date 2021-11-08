package Commands

import (
	"fmt"
	"ObjectOrientedCourse/src/Rooms"
)

func CommandTalk(room *Rooms.Room) {
	room.GetCharacter().Talk()
}

func CommandMove(room *Rooms.Room) *Rooms.Room {
	var direction string
	for {
		fmt.Println("Enter direction")
		fmt.Scanln(&direction)
		newRoom := room.Move(direction)
		if newRoom != room {
			fmt.Println("")
			return newRoom
		}
	}
}