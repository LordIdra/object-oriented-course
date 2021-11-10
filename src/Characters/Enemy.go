package Characters

import (
	"fmt"
	"strconv"
)

type Enemy struct {
	Neutral
	health  int
}

func NewEnemy(name string, description string, conversation string, health int) *Enemy {
	enemy := new(Enemy)
	enemy.name = name
	enemy.description = description
	enemy.conversation = conversation
	enemy.health = health
	return enemy
}

func formatHealth(health int, maxHealth int) string {
	bar := ""
	for i := 0; i < 20; i++ {
		if float32(i) < (float32(health) / float32(maxHealth) * 20) {
			break
		}
		bar += "|"
	}
	return bar
} 

func (enemy *Enemy) Fight(health *int, maxHealth int) bool {
	fmt.Println("----- FIGHT -----")
	fmt.Println("You: [" + formatHealth(*health, maxHealth) + "] " + strconv.Itoa(*health) + "/" + strconv.Itoa(maxHealth))
	return true;
}