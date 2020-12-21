package bingo

import (
	"fmt"
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

func (p *Player) CheckNumber(number int) {
	for idx := range p.Cards {
		card := &p.Cards[idx]
		card.checkNumber(number)
	}
}
