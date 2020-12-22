package bingo

// Player user type
type Game struct {
	ID   int
	Name string
}

// NewPlayer creates a new player
func NewGame() Game {
	// id :=
	p := Game{
		ID:   0,
		Name: "",
	}
	return p
}
