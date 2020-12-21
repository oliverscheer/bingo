package bingo

import (
	"fmt"
	"math/rand"
)

// BingoCard defines a bingo card
type BingoCard struct {
	ID     int
	Cell   [5][5]BingoCardCell
	Player Player
}

// CheckNumber checks the number on a card
func (b *BingoCard) checkNumber(number int) bool {
	win := false
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			cell := &b.Cell[x][y]
			if cell.Value == number {
				cell.Hit = true
				// TODO: Fire Hit event
				win = b.checkForWin()
				if win {
					break
				}
			}
		}
	}
	return win
}

func (b *BingoCard) checkForWin() bool {
	win := false

	// Vertical line check
	for x := 0; x < 5; x++ {
		counter := 0
		for y := 0; y < 5; y++ {
			if b.Cell[x][y].Hit {
				counter++
			}
		}
		if counter == 5 {
			win = true
		}
	}

	// Horizontal line check
	for y := 0; y < 5; y++ {
		counter := 0
		for x := 0; x < 5; x++ {
			if b.Cell[x][y].Hit {
				counter++
			}
		}
		if counter == 5 {
			win = true
		}
	}

	// cross 1 check
	counter := 0
	for x := 0; x < 5; x++ {
		if x == 2 {
			counter++
			continue
		}
		if b.Cell[x][x].Hit {
			counter++
		}
		if counter == 5 {
			win = true
		}
	}

	// cross 2 check
	counter = 0
	for x := 0; x < 5; x++ {
		y := 4 - x
		if x == 2 {
			counter++
			continue
		}
		if b.Cell[x][y].Hit {
			counter++
		}
		if counter == 5 {
			win = true
		}
	}

	if win {
		fmt.Println()
		fmt.Println(b.Player.Name, "wins")
		b.PrintDetails()
		return true
	}

	return false
}

// PrintDetails show the details of a card
func (b *BingoCard) PrintDetails() {
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
func CreateCards(numberOfCards int) []BingoCard {
	var cardList = []BingoCard{}
	for currentCardCount := 0; currentCardCount < numberOfCards; currentCardCount++ {
		card := BingoCard{
			ID: currentCardCount,
		}
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				cell := &card.Cell[x][y]
				if (x == 2) && (y == 2) {
					cell.Hit = true
					continue
				}
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
