package bingo

import (
	"fmt"
	"strconv"
)

// Player user type
type Player struct {
	ID    int
	Name  string
	Cards []BingoCard
}

// NewPlayer creates a new player
func NewPlayer(id int, name string) Player {
	p := Player{ID: id, Name: name}
	return p
}

// AddCard adds a card to a player
func (p *Player) AddCard(card BingoCard) {
	card.Player = *p
	cards := append(p.Cards, card)
	p.Cards = cards
}

func (p *Player) printDetails() {
	fmt.Println("\nPlayer:", p.Name)
	for idx := range p.Cards {
		card := &p.Cards[idx]
		card.PrintDetails()
	}
}

func (p *Player) CheckNumber(number int) bool {
	win := false
	for idx := range p.Cards {
		card := &p.Cards[idx]
		win = card.checkNumber(number)
		if win {
			break
		}
	}
	return win
}

func CreatePlayer(numberOfPlayers int) []Player {

	// Create player
	var playerList = []Player{}
	for currentPlayerNo := 0; currentPlayerNo < numberOfPlayers; currentPlayerNo++ {
		var playerName string = "Player " + strconv.Itoa(currentPlayerNo)
		newPlayer := NewPlayer(currentPlayerNo, playerName)
		playerList = append(playerList, newPlayer)
	}
	fmt.Println("Generated", numberOfPlayers, "player")

	return playerList
}
