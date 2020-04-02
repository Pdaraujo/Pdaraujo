package main

import (
	"defaults"
	"fmt"
	"solver"
	"time"
)

func main() {
	game := solver.NewGame(defaults.DefaultGame)
	game.Print()
	initTime := time.Now()
	game.Solve()
	fmt.Println("\n\n")
	game.Print()
	fmt.Printf("\n\nIterations: %d", game.Iteration)
	fmt.Printf("\nIn total time: %s", time.Now().Sub(initTime))
}

//This was just a test
func testingChannels(game * solver.Game) {
	c := make(chan *solver.Position)
	game.FindFirstEmpty(c)

	for position := range c {
		if position != nil {
			fmt.Printf("\n\nRow: %d\nColumn: %d", position.Row, position.Col)
		}
	}
}