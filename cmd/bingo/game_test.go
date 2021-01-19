package bingo

import "testing"

func TestCreateAndStartGameWithFivePlayer(t *testing.T) {

	settings := GameSettings{
		Delay:   100,
		Cards:   2,
		Verbose: true,
	}
	game := NewGame(settings)

	player := NewPlayer(1, "Oliver")
	game.AddPlayer(player)

	player = NewPlayer(2, "Johanna")
	game.AddPlayer(player)

	player = NewPlayer(3, "Antonia")
	game.AddPlayer(player)

	player = NewPlayer(4, "Valentina")
	game.AddPlayer(player)

	player = NewPlayer(5, "Cosima")
	game.AddPlayer(player)

	gameData := game.Play()

	playerCount := len(game.GetPlayer())
	wanted := 5
	if playerCount != wanted {
		t.Errorf("Wrong number of player, wanted %d, got %d", wanted, playerCount)
	}

	if gameData.Rounds == 0 {
		t.Errorf("No rounds played. got: %d", gameData.Rounds)
	}
}
