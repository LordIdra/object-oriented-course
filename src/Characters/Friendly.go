package Characters

import (
	"ObjectOrientedCourse/src/Items"
	"fmt"
)

type Friendly struct {
	Neutral
	item Items.Item
}

func NewFriendly(name string, description string, conversation string, item Items.Item) *Friendly {
	friendly := new(Friendly)
	friendly.name = name
	friendly.description = description
	friendly.conversation = conversation
	friendly.item = item
	return friendly
}

func (friendly *Friendly) Talk() {
	fmt.Println(friendly.conversation)
}

func (friendly *Friendly) GetItem() Items.Item {
	return friendly.item
}