package Rooms

import "ObjectOrientedCourse/src/Characters"

var rooms []*Room

func createRooms() {
	rooms = append(rooms, NewRoom("Kitchen", "The knives are located here"))
	rooms = append(rooms, NewRoom("Dining Hall", "Exists"))
	rooms = append(rooms, NewRoom("Ballroom", "Stay away from here, I'm bad at dancing"))
}

func linkRooms() {
	rooms[0].addLink("south", rooms[1])
	rooms[1].addLink("north", rooms[0])
	rooms[2].addLink("east", rooms[1])
	rooms[1].addLink("west", rooms[2])
}

func createCharacters() {
	rooms[0].setCharacter(Characters.NewCharacter(
		"Dave",
		"BRAAAAAAIIIIINNNNSSSS",
		"Hello little one"))
}

func Initialize() {
	createRooms()
	linkRooms()
	createCharacters()
}
