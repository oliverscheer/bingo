package main

import (
	"flag"
	"fmt"
	"math/rand"
	"my/bingo"
	"time"
)

// defaults
const defaultNumberOfPlayer int = 1000
const defaultNumberOfCardsForEachPlayer int = 1000
const defaultSleepInMS int = 500
const defaultIterations int = 1

// TODO: Everything is linear. Always the first player in the row wins with the right numbers

// TODO:

func main() {

	// commandline args
	numberOfPlayers := flag.Int("player", defaultNumberOfPlayer, "number of simulated player")
	numberOfCardsForEachPlayer := flag.Int("cards", defaultNumberOfCardsForEachPlayer, "number of cards for each player")
	sleepInMS := flag.Int("sleep", defaultSleepInMS, "delay for each round in milliseconds")
	iterations := flag.Int("iterations", defaultIterations, "number of interations")
	// var svar string
	// flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	// TODO: Add version number here
	fmt.Println("BinGO v.0.1")

	fmt.Println("- player    :", *numberOfPlayers)
	fmt.Println("- cards     :", *numberOfCardsForEachPlayer)
	fmt.Println("- sleep     :", *sleepInMS)
	fmt.Println("- iterations:", *iterations)
	fmt.Println()

	// fmt.Println("tail:", flag.Args())

	for iteration := 1; iteration <= *iterations; iteration++ {

		fmt.Println("### Iteration", iteration, "####")
		// Initialize real random values
		rand.Seed(time.Now().UnixNano())

		// Create player
		var playerList = bingo.CreatePlayer(*numberOfPlayers)

		// Create cards
		var numberOfCards = *numberOfPlayers * *numberOfCardsForEachPlayer
		var cardList = bingo.CreateCards(numberOfCards)

		// Distribute cards to player
		currentCardID := 0
		newCurrentCardID := 0
		for idx := range playerList {
			player := &playerList[idx]
			for cardID := currentCardID; cardID < (currentCardID + *numberOfCardsForEachPlayer); {
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

		// linear way: get number and send it to player
		win := false
		for round := 1; round <= 75; round++ {
			ms := time.Duration(*sleepInMS)
			time.Sleep(ms * time.Millisecond)
			currentNumber := pot[round-1]
			fmt.Println("Round: ", round, " - Number: ", currentNumber)
			for idx := range playerList {
				player := &playerList[idx]
				win = player.CheckNumber(currentNumber)
				if win {
					break
				}
			}
			if win {
				break
			}
		}

		// parallel way: get number and send it to player
		// for round := 1; round <= 75; round++ {
		// 	time.Sleep(sleepInMS * time.Millisecond)
		// 	currentNumber := pot[round-1]
		// 	fmt.Println("Round: ", round, " - Number: ", currentNumber)
		// 	for idx := range playerList {
		// 		player := &playerList[idx]
		// 		player.CheckNumber(currentNumber)
		// 	}
		// }
	}
}
