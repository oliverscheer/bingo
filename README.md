# Bingo

A Golang training project to play binGO. 

## bingo library

```bash
cd cmd/bingo

# run tests
go test

# run fresh source
go run main.app

# compile bingocli
go build .

# run the compiled version with parameters
./bingocli -player 5 -iterations 5 -cards 5 -sleep 100
```

## bingocli-app - Basic Stand-alone Bingo Simulation 

```bash
cd cmd/app

# run tests
go test

# run fresh source
go run main.app

# compile bingocli
go build .

# run the compiled version with parameters
./bingocli -player 5 -iterations 5 -cards 5 -sleep 100
```

## webserver

```bash
cd cmd/webserver

# run tests
go test

# Start webserver
go run main.app

# Browse to http://localhost:8080/alive to check if server is running

```

## To do's

1. Everything is still linear. Always the first player in the row wins with the right numbers
Needs to be paralyzed.

1. Add time tracking for performance measuring

1. Add linter for markdown 