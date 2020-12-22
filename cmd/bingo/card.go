package bingo

import (
	"fmt"
	"math/rand"
)

// Card defines a bingo card
type Card struct {
	ID     int
	Cell   [5][5]CardCell
	Player Player
}

// CheckNumber checks the number on a card
func (b *Card) CheckNumber(number int) bool {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			cell := &b.Cell[x][y]
			hit := cell.CheckNumber(number)
			if hit {
				return true
			}
		}
	}
	return false
}

// CheckForWin checks the card for a win
func (b *Card) CheckForWin() bool {
	// Vertical line check
	for x := 0; x < 5; x++ {
		counter := 0
		for y := 0; y < 5; y++ {
			cell := b.Cell[x][y]
			if cell.IsHit() {
				counter++
			}
		}
		if counter == 5 {
			return true
		}
	}

	// Horizontal line check
	for y := 0; y < 5; y++ {
		counter := 0
		for x := 0; x < 5; x++ {
			cell := b.Cell[x][y]
			if cell.IsHit() {
				counter++
			}
		}
		if counter == 5 {
			return true
		}
	}

	// cross 1 check
	counter := 0
	for x := 0; x < 5; x++ {
		cell := b.Cell[x][x]
		if cell.IsHit() {
			counter++
		}
		if counter == 5 {
			return true
		}
	}

	// cross 2 check
	counter = 0
	for x := 0; x < 5; x++ {
		y := 4 - x
		cell := b.Cell[x][y]
		if cell.IsHit() {
			counter++
		}
		if counter == 5 {
			return true
		}
	}

	return false
}

// PrintDetails show the details of a card
func (b *Card) PrintDetails() {
	fmt.Println()
	fmt.Println("| CardID: ", b.ID)
	fmt.Println("|  B   |  I   |  N   |  G   |  O   |")
	fmt.Println("|------|------|------|------|------|")
	for y := 0; y < 5; y++ {
		fmt.Print("|")
		for x := 0; x < 5; x++ {
			b.Cell[x][y].printDetails()
			fmt.Print(" |")
		}
		fmt.Println()
		fmt.Println("|------|------|------|------|------|")
	}
}

// CreateCards generates a specific number of unique cards
func CreateCards(numberOfCards int) []Card {
	var cardList = []Card{}
	for currentCardCount := 0; currentCardCount < numberOfCards; currentCardCount++ {
		card := Card{
			ID: currentCardCount,
		}
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				cell := &card.Cell[x][y]

				var v int
				v = rand.Intn(15) + 1 + x*15
				if y > 0 {
					var check bool = false
					for !check {
						v = rand.Intn(15) + x*15 + 1
						check = true
						for ycheck := 0; ycheck < y; ycheck++ {
							if card.Cell[x][ycheck].Value == v {
								// number already exist, we need a new one
								check = false
								break
							}
						}
					}
				}
				cell.Value = v
			}
		}

		// TODO: I do not check if similar card already exist
		cardList = append(cardList, card)
	}
	fmt.Println("Generated", numberOfCards, "bingo cards")
	return cardList
}
