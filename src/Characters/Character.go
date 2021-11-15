package Characters

import "ObjectOrientedCourse/src/Items"

type Character interface {
	Describe()
	Talk()
	GetItem() Items.Item
	Fight(health *int, maxHealth int, items []Items.Item) bool
}
