package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// TODO: Everything is linear. Always the first player in the row wins with the right numbers

func main() {
	// TODO: Add version number here
	fmt.Println("BinGO v.0.1")

	const numberOfCardsForEachPlayer = 1000
	const numberOfPlayers = 1000
	const sleepInMS = 500

	rand.Seed(time.Now().UnixNano())

	// Create player
	var playerList = []Player{}
	for currentPlayerNo := 0; currentPlayerNo < numberOfPlayers; currentPlayerNo++ {
		var playerName string = "Player " + strconv.Itoa(currentPlayerNo)
		newPlayer := NewPlayer(currentPlayerNo, playerName)
		playerList = append(playerList, newPlayer)
	}
	fmt.Println("Generated", numberOfPlayers, "player")

	// Create cards
	var numberOfCards = numberOfPlayers * numberOfCardsForEachPlayer
	var cardList = []BingoCard{}
	for currentCardCount := 0; currentCardCount < numberOfCards; currentCardCount++ {
		card := BingoCard{
			ID: currentCardCount,
		}
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				cell := &card.Cell[x][y]
				if (x == 2) && (y == 2) {
					cell.hit = true
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
							if card.Cell[x][ycheck].value == v {
								// number already exist, we need a new one
								check = false
								break
							}
						}
					}
				}
				cell.value = v
			}
		}

		// TODO: I do not check if similar card already exist
		cardList = append(cardList, card)
	}
	fmt.Println("Generated", numberOfCards, "bingo cards")

	// for idx := range cardList {
	// 	card := cardList[idx]
	// 	card.PrintDetails()
	// }
	// // return

	// Distribute cards to player
	currentCardID := 0
	newCurrentCardID := 0
	for idx := range playerList {
		player := &playerList[idx]
		for cardID := currentCardID; cardID < (currentCardID + numberOfCardsForEachPlayer); {
			player.AddCard(cardList[cardID])
			cardID++
			newCurrentCardID = cardID
		}
		currentCardID = newCurrentCardID
	}

	// create pot
	var pot []int
	for ball := 0; ball < 75; ball++ {
		var alreadyInPot bool = true
		var newNumber int
		for alreadyInPot {
			newNumber = rand.Intn(75) + 1
			if ball == 0 {
				alreadyInPot = false
				break
			}
			var found bool = false
			for i := 0; i < ball; i++ {
				if pot[i] == newNumber {
					found = true
					break
				}
			}
			if !found {
				alreadyInPot = false
			}
		}
		pot = append(pot, ball)
		pot[ball] = newNumber
	}

	// get number and send it to player
	for round := 1; round <= 75; round++ {
		time.Sleep(sleepInMS * time.Millisecond)
		currentNumber := pot[round-1]
		fmt.Println("Round: ", round, " - Number: ", currentNumber)
		for idx := range playerList {
			player := &playerList[idx]
			player.checkNumber(currentNumber)
		}
	}
}

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

func (p *Player) checkNumber(number int) {
	for idx := range p.Cards {
		card := &p.Cards[idx]
		card.checkNumber(number)
	}
}

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
			if cell.value == number {
				cell.hit = true
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
			if b.Cell[x][y].hit {
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
			if b.Cell[x][y].hit {
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
		if b.Cell[x][x].hit {
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
		if b.Cell[x][y].hit {
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

// BingoCardCell describes a BingoCard cell
type BingoCardCell struct {
	value int
	hit   bool
}

func (bc *BingoCardCell) printDetails() {
	var v int
	v = bc.value
	var s string
	if v == 0 {
		s = " *"
	} else if v < 10 {
		s = " " + strconv.Itoa(v)
	} else {
		s = strconv.Itoa(v)
	}
	if bc.hit {
		s = "*" + s + "*"
	} else {
		s = " " + s + " "
	}
	fmt.Print(" ", s)
}
