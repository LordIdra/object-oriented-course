package Rooms

import (
	"ObjectOrientedCourse/src/Characters"
	"ObjectOrientedCourse/src/Items"
)

var rooms []*Room

func createRooms() {
	rooms = append(rooms, NewRoom("Kitchen", "The knives are located here"))
	rooms = append(rooms, NewRoom("Dining Hall", "Exists"))
	rooms = append(rooms, NewRoom("Ballroom", "Stay away from here, I'm bad at dancing"))
	rooms = append(rooms, NewRoom("Living Room", "Just chillin'"))
}

func linkRooms() {
	rooms[0].addLink("south", rooms[1])
	rooms[1].addLink("north", rooms[0])
	rooms[2].addLink("east", rooms[1])
	rooms[1].addLink("west", rooms[2])
	rooms[3].addLink("west", rooms[0])
	rooms[0].addLink("east", rooms[3])
}

func createCharacters() {
	rooms[0].setCharacter(Characters.NewEnemy(
		"Dave",
		"BRAAAAAAIIIIINNNNSSSS",
		"Hello little one",
		50, 2, 4))
	rooms[1].setCharacter(Characters.NewFriendly(
		"Bob",
		"A builder.",
		"It's dangerous to go alone! Take this wrench.",
		Items.Item{
			Name: "Wrench", 
			MinDamage: 10, 
			MaxDamage: 15}))
}

func Initialize() {
	createRooms()
	linkRooms()
	createCharacters()
}

func DefaultRoom() *Room {
	return rooms[0]
}