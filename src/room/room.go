package room

import "ObjectOrientedCourse/src/character"


type Room struct {
	name string
	description string
	links map[string]*Room
	character Character
}


func NewRoom(name string, description string) *Room {
	room := new(Room)
	room.name = name
	room.description = description
	return room
}

func Name(room Room) string {
	return room.name
}