package bingo

// User user type
type Player struct {
	ID   int64
	Name string
	Card *BingoCard
}

// Address address type
type BingoCard struct {
	City   string
	ZIP    int
	LatLng [2]float64
}

var player = Player{}

// Hello writes a welcome string
func Hello() string {
	return "Hello, " + player.Name
}
