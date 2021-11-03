package Commands

import (
	"ObjectOrientedCourse/src/Rooms"
)

func CommandTalk(room *Rooms.Room) {
	room.GetCharacter().Talk()
}