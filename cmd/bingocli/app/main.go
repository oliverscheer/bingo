package main

import (
	"fmt"
	"math/rand"
	"my/bingo"
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
	var playerList = []bingo.Player{}
	for currentPlayerNo := 0; currentPlayerNo < numberOfPlayers; currentPlayerNo++ {
		var playerName string = "Player " + strconv.Itoa(currentPlayerNo)
		newPlayer := bingo.NewPlayer(currentPlayerNo, playerName)
		playerList = append(playerList, newPlayer)
	}
	fmt.Println("Generated", numberOfPlayers, "player")

	// Create cards
	var numberOfCards = numberOfPlayers * numberOfCardsForEachPlayer
	var cardList = []bingo.BingoCard{}
	for currentCardCount := 0; currentCardCount < numberOfCards; currentCardCount++ {
		card := bingo.BingoCard{
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
			player.CheckNumber(currentNumber)
		}
	}
}
