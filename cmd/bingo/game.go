package bingo

import (
	"fmt"
	"math/rand"
	"time"
)

// Game user type
type Game struct {
	ID         int
	Name       string
	playerList []Player
	settings   GameSettings
}

type GameData struct {
	Rounds     int
	Winner     Player
	WinnerCard Card
}

type GameSettings struct {
	Delay   int
	Cards   int
	Verbose bool
}

// NewGame creates a new game
func NewGame(settings GameSettings) Game {
	// id :=
	p := Game{
		ID:       0,
		Name:     "",
		settings: settings,
	}

	return p
}

// AddPlayer adds a player to the game
func (g *Game) AddPlayer(player Player) {
	g.playerList = append(g.playerList, player)
}

// GetPlayer gets current player of the game
func (g *Game) GetPlayer() []Player {
	return g.playerList
}

// Play starts the game
func (g *Game) Play() GameData {
	gameData := GameData{}

	if g.settings.Verbose {
		fmt.Println("### Game:", g.ID, "####")
	}

	// Initialize real random values
	rand.Seed(time.Now().UnixNano())

	// // Create player
	// var playerList = bingo.CreatePlayer(*numberOfPlayers)

	// var game := bingo.NewGame(iteration)
	// for p, _ := range &playerList {
	// 	game.addPlayer(p)
	// }

	playerCount := len(g.GetPlayer())

	// Create cards
	var numberOfCards = playerCount * g.settings.Cards
	var cardList = CreateCards(numberOfCards)

	// Distribute cards to player
	currentCardID := 0
	newCurrentCardID := 0
	for idx := range g.playerList {
		player := &g.playerList[idx]
		for cardID := currentCardID; cardID < (currentCardID + g.settings.Cards); {
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
		gameData.Rounds = round
		ms := time.Duration(g.settings.Delay)
		time.Sleep(ms * time.Millisecond)
		currentNumber := pot[round-1]
		fmt.Println("Round: ", round, " - Number: ", currentNumber)
		for idx := range g.playerList {
			player := &g.playerList[idx]
			hit := player.CheckNumber(currentNumber)
			if hit {
				if g.settings.Verbose {
					fmt.Println(player.Name, "hit")
				}
			}
			win = player.CheckForWin()
			if win {
				if g.settings.Verbose {
					fmt.Println(player.Name, "wins")
				}
				gameData.Winner = *player
				break
			}
		}
		if win {
			break
		}
	}

	return gameData
}
