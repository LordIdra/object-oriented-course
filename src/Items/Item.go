package Items

import (
	"fmt"
	"strconv"
)

type Item struct {
	Name      string
	MinDamage int
	MaxDamage int
}

func (item *Item) Describe() {
	fmt.Println("- " + item.Name + " (" + strconv.Itoa(item.MinDamage) + "-" + strconv.Itoa(item.MaxDamage) + ")")
}
