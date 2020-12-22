package main

import (
	"bingo"
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// defaults
const defaultNumberOfPlayer int = 1000
const defaultNumberOfCardsForEachPlayer int = 1000
const defaultSleepInMS int = 500
const defaultIterations int = 1

// TODO: Everything is linear. Always the first player in the row wins with the right numbers

// Config defines settings for the game
type Config struct {
	verbose                    bool
	numberOfPlayers            int
	numberOfCardsForEachPlayer int
	sleepInMS                  int
	iterations                 int

	// args are the positional (non-flag) command-line arguments.
	args []string
}

// parseFlags parses the command-line arguments provided to the program.
// Typically os.Args[0] is provided as 'progname' and os.args[1:] as 'args'.
// Returns the Config in case parsing succeeded, or an error. In any case, the
// output of the flag.Parse is returned in output.
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message.
func parseFlags(progname string, args []string) (config *Config, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf Config
	flags.BoolVar(&conf.verbose, "verbose", false, "set verbosity")
	flags.IntVar(&conf.numberOfPlayers, "player", defaultNumberOfPlayer, "number of simulated player")
	flags.IntVar(&conf.numberOfCardsForEachPlayer, "cards", defaultNumberOfCardsForEachPlayer, "number of cards for each player")
	flags.IntVar(&conf.sleepInMS, "sleep", defaultSleepInMS, "delay for each round in milliseconds")
	flags.IntVar(&conf.iterations, "iterations", defaultIterations, "number of interations")

	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	conf.args = flags.Args()
	return &conf, buf.String(), nil
}

func startGame(config Config) {
	for iteration := 1; iteration <= config.iterations; iteration++ {
		fmt.Println("### Iteration", iteration, "####")

		// Initialize real random values
		rand.Seed(time.Now().UnixNano())

		// Create player
		var playerList = bingo.CreatePlayer(config.numberOfPlayers)

		settings := bingo.GameSettings{
			Cards: config.numberOfCardsForEachPlayer,
			Delay: config.sleepInMS,
		}

		var game = bingo.NewGame(settings)
		for _, player := range playerList {
			game.AddPlayer(player)
		}

		gameData := game.Play()

		fmt.Println(gameData.Winner.Name, "wins.")
	}

}

func main() {

	// checking args
	config, output, err := parseFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("got error:", err)
		fmt.Println("output:\n", output)
		os.Exit(1)
	}

	fmt.Println("BinGO v.0.1")
	fmt.Println("- player    :", config.numberOfPlayers)
	fmt.Println("- cards     :", config.numberOfCardsForEachPlayer)
	fmt.Println("- sleep     :", config.sleepInMS)
	fmt.Println("- iterations:", config.iterations)
	fmt.Println()

	startGame(*config)
}
