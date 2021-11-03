package Rooms

import (
	"ObjectOrientedCourse/src/Characters"
	"fmt"
)

type Room struct {
	name        string
	description string
	links       map[string]*Room
	character   *Characters.Character
}

func NewRoom(name string, description string) *Room {
	room := new(Room)
	room.name = name
	room.description = description
	return room
}

func (room *Room) addLink(direction string, link *Room) {
	room.links[direction] = link
}

func (room *Room) setCharacter(character *Characters.Character) {
	room.character = character
}

func (room *Room) GetCharacter() *Characters.Character {
	return room.character
}

func (room *Room) Describe() {
	fmt.Println("--- ", room.name, " ---")
	fmt.Println(room.description)
	for direction, room := range room.links {
		fmt.Println(room, " is ", direction)
	}
	if room.character != nil {
		room.character.Describe()
	}
}

func (room *Room) Move(direction string) *Room {
	if room.links[direction] == nil {
		fmt.Println("Invalid direction!")
		return room
	} else {
		return room.links[direction]
	}
}
