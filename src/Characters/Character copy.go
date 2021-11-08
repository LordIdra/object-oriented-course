package Characters

type Enemy struct {
	character Character
	health    int
}

func NewEnemy(name string, description string, conversation string, health int) *Enemy {
	character := NewCharacter(name, description, conversation)
	enemy := new(Enemy)
	enemy.character = *character
	enemy.health = health
	return enemy
}
