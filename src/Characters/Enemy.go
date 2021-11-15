package Characters

import (
	"ObjectOrientedCourse/src/Items"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Enemy struct {
	Neutral
	health    int
	maxHealth int
	minDamage int
	maxDamage int
}

func NewEnemy(name string, description string, conversation string, health int, minDamage int, maxDamage int) *Enemy {
	enemy := new(Enemy)
	enemy.name = name
	enemy.description = description
	enemy.conversation = conversation
	enemy.health = health
	enemy.maxHealth = health
	enemy.minDamage = minDamage
	enemy.maxDamage = maxDamage
	return enemy
}

func formatHealth(health int, maxHealth int) string {
	bar := ""
	for i := 0; i < 20; i++ {
		if float32(i) < (float32(health) * 20 / float32(maxHealth)) {
			bar += "|"
		} else {
			bar += " "
		}
	}
	return bar
}

func printHealthAndItems(health int, maxHealth int, enemy *Enemy, items []Items.Item) {
	fmt.Println("")
	fmt.Println("You:   [" + formatHealth(health, maxHealth) + "] " + strconv.Itoa(health) + "/" + strconv.Itoa(maxHealth))
	fmt.Println("Enemy: [" + formatHealth(enemy.health, enemy.maxHealth) + "] " + strconv.Itoa(enemy.health) + "/" + strconv.Itoa(enemy.maxHealth))
	fmt.Println("Your Items:")
	for _, item := range items {
		item.Describe()
	}
}

func getItemToFightWith(items []Items.Item) Items.Item {
	var itemName string
	for {
		fmt.Println("What will you fight with?")
		fmt.Scanln(&itemName)
		for _, item := range items {
			if strings.EqualFold(item.Name, itemName) {
				return item
			}
		}
		fmt.Println("Invalid item!")
	}
}

func randomNumber(n int) int {
	randomSource := rand.NewSource(int64(time.Now().Nanosecond()))
	randomGenerator := rand.New(randomSource)
	return randomGenerator.Intn(n)
}

func (enemy *Enemy) damagePlayer(health *int) {
	damage := enemy.minDamage + randomNumber(enemy.maxDamage-enemy.minDamage)
	*health -= damage
	fmt.Println("The enemy does " + strconv.Itoa(damage) + " damage.")
}

func (enemy *Enemy) applyDamage(item Items.Item) {
	damage := item.MinDamage + randomNumber(item.MaxDamage-item.MinDamage)
	enemy.health -= damage
	fmt.Println("You do " + strconv.Itoa(damage) + " damage.")
}

func (enemy *Enemy) Fight(health *int, maxHealth int, items []Items.Item) bool {
	for {
		printHealthAndItems(*health, maxHealth, enemy, items)
		item := getItemToFightWith(items)
		enemy.applyDamage(item)
		enemy.damagePlayer(health)
		if enemy.health < 0 {
			fmt.Println("------ YOU WIN ------")
			fmt.Println("")
			enemy = nil
			return true
		} else if *health <= 0 {
			fmt.Println("----- GAME OVER -----")
			fmt.Println("")
			return false
		}
	}
}
