package Commands

import (
	"ObjectOrientedCourse/src/Items"
	"ObjectOrientedCourse/src/Rooms"
	"fmt"
)

func CommandTalk(room *Rooms.Room) Items.Item {
	room.GetCharacter().Talk()
	return room.GetCharacter().GetItem()
}

func CommandFight(room *Rooms.Room, health *int, maxHealth int, items []Items.Item) bool {
	fmt.Println("----- FIGHT -----")
	return room.GetCharacter().Fight(health, maxHealth, items)
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