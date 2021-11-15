package Characters

import (
	"ObjectOrientedCourse/src/Items"
	"fmt"
)

type Neutral struct {
	name         string
	description  string
	conversation string
}

func NewNeutral(name string, description string, conversation string) *Neutral {
	neutral := new(Neutral)
	neutral.name = name
	neutral.description = description
	neutral.conversation = conversation
	return neutral
}

func (neutral *Neutral) Describe() {
	fmt.Println("")
	fmt.Println("[" + neutral.name + "] " + neutral.description)
}

func (neutral *Neutral) Talk() {
	if neutral.conversation == "" {
		fmt.Println(neutral.name + " doesn't want to talk to you.")
	} else {
		fmt.Println(neutral.name + " says: " + neutral.conversation)
	}
}

func (neutral *Neutral) Fight(health *int, maxHealth int, items []Items.Item) bool {
	fmt.Println(neutral.name + " doesn't want to fight with you.")
	return true
}

func (neutral *Neutral) GetItem() Items.Item {
	return Items.Item{}
}