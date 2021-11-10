package Characters

type Character interface {
	Describe()
	Talk()
	Fight(health *int, maxHealth int) bool
}
