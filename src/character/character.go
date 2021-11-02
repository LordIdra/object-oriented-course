package character


type Character struct {
	name string
	description string
	conversation string
}


func NewCharacter(name string, description string, conversation string) *Character{
	character := new(Character)
	character.name = name
	character.description = description
	character.conversation = conversation
	return character
}

func Name(character Character) string {
	return character.name
}

func Description(character Character) string {
	return character.description
}

func Conversation(character Character) string {
	return character.conversation
}