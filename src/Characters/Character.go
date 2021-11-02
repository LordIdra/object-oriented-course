package Characters

import "fmt"

type Character struct {
	name         string
	description  string
	conversation string
}

func NewCharacter(name string, description string, conversation string) *Character {
	character := new(Character)
	character.name = name
	character.description = description
	character.conversation = conversation
	return character
}

func (character *Character) Describe() {
	fmt.Print("")
	fmt.Print("[", character.name, "] ", character.description)
}

func (character *Character) Talk() {
	if character.conversation == "" {
		fmt.Print("[", character.name, "] ", "Doesn't want to talk to you.")
	} else {
		fmt.Print("[", character.name, "] ", character.conversation)
	}
}
