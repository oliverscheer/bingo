package bingo

import (
	"fmt"
	"strconv"
)

// Player user type
type Player struct {
	ID    int
	Name  string
	Cards []Card
}

// NewPlayer creates a new player
func NewPlayer(id int, name string) Player {
	p := Player{
		ID:   id,
		Name: name,
	}
	return p
}

// AddCard adds a card to a player
func (p *Player) AddCard(card Card) {
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

// CheckNumber checks if card contains card
func (p *Player) CheckNumber(number int) bool {
	for idx := range p.Cards {
		card := &p.Cards[idx]
		if card.CheckNumber(number) {
			return true
		}
	}
	return false
}

func (p *Player) CheckForWin() bool {
	for idx := range p.Cards {
		card := &p.Cards[idx]
		if card.CheckForWin() {
			return true
		}
	}
	return false
}

// CreatePlayer creates an array of player
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
