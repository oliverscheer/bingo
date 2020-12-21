package bingo

import (
	"fmt"
	"os"
)

// BingoCard defines a bingo card
type BingoCard struct {
	ID     int
	Cell   [5][5]BingoCardCell
	Player Player
}

// CheckNumber checks the number on a card
func (b *BingoCard) checkNumber(number int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			cell := &b.Cell[x][y]
			if cell.Value == number {
				cell.Hit = true
				// TODO: Fire Hit event
				b.checkForWin()
				break
			}
		}
	}
}

func (b *BingoCard) checkForWin() {
	// Vertical line check
	for x := 0; x < 5; x++ {
		counter := 0
		for y := 0; y < 5; y++ {
			if b.Cell[x][y].Hit {
				counter++
			}
		}
		if counter == 5 {
			// TODO: Fire win event
			fmt.Println()
			fmt.Println(b.Player.Name, "wins")
			b.PrintDetails()
			os.Exit(0)
			return
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
			// TODO: Fire win event
			fmt.Println()
			fmt.Println(b.Player.Name, "wins")
			b.PrintDetails()
			os.Exit(0)
			return
		}
	}

	// cross 1 check
	for x := 0; x < 5; x++ {
		counter := 0
		if b.Cell[x][x].Hit {
			counter++
		}
		if counter == 5 {
			// TODO: Fire win event
			fmt.Println()
			fmt.Println(b.Player.Name, "wins")
			b.PrintDetails()
			os.Exit(0)
			return
		}
	}

	// cross 2 check
	for x := 0; x < 5; x++ {
		y := 4 - x
		counter := 0
		if b.Cell[x][y].Hit {
			counter++
		}
		if counter == 5 {
			// TODO: Fire win event
			fmt.Println()
			fmt.Println(b.Player.Name, "wins")
			b.PrintDetails()
			os.Exit(0)
			return
		}
	}
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
