package solver

import (
	"errors"
	"fmt"
	"math"
)

type GameLogic interface {
	Print()
	FindEmpty() (*Position, error)
	Valid(int, *Position) bool
	Solve() bool
}

type Position struct {
	Row int
	Col int
}

type Board [9][9]int

type Game struct {
	Board *Board
	Iteration int
}

func NewGame(board [9][9]int) *Game {
	var new Board = board
	return &Game{Board: &new}
}

func (g * Game) Print() {
	for row := range g.Board {
		if row % 3 == 0 && row != 0 {
			fmt.Println("\n+---------+-----------+---------+")
		} else {
			fmt.Println()
		}

		for col := range g.Board[0] {
			if col % 3 == 0 && col != 0 {
				fmt.Print(" | ")
			}

			if col == len(g.Board[0]) - 1 {
				fmt.Printf(" %d", g.Board[row][col])
			} else {
				fmt.Printf(" %d ", g.Board[row][col])
			}
		}
	}
}

func (g * Game) FindEmpty() (*Position, error) {
	for row := range g.Board {
		for col := range g.Board[0] {
			if g.Board[row][col] == 0 {
				return &Position{row, col}, nil
			}
		}
	}
	return nil, errors.New("no empty position found")
}

func (g *Game) Valid(num int, pos *Position) bool {
	//Check Row
	for c := range g.Board[0] {
		if g.Board[pos.Row][c] == num && c != pos.Col {
			return false
		}
	}

	//Check column
	for r := range g.Board {
		if g.Board[r][pos.Col] == num && r != pos.Row {
			return false
		}
	}

	//Check Box
	boxX, boxY := getBoxFor(pos)

	for _, r := range makeRange(boxY * 3, boxY * 3 + 3) {
		for _, c := range makeRange(boxX * 3, boxX * 3 + 3) {
			if g.Board[r][c] == num && pos.Row != r && pos.Col != c {
				return false
			}
		}
	}
	return true
}

func (g *Game) Solve() bool {
	g.Iteration += 1
	emptyPosition, err := g.FindEmpty()

	if err != nil {
		return true
	}

	for _, num := range makeRange(1, 10) {
		if g.Valid(num, emptyPosition) {
			g.Board[emptyPosition.Row][emptyPosition.Col] = num

			if g.Solve() {
				return true
			} else {
				g.Board[emptyPosition.Row][emptyPosition.Col] = 0
			}
		}
	}
	return false
}

//Helpers
func getBoxFor(pos *Position) (int, int) {
	var boxX int = int(math.Floor(float64(pos.Col / 3)))
	var boxY int = int(math.Floor(float64(pos.Row / 3)))

	return boxX, boxY
}

func makeRange(min, max int) []int {
	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}




func (g *Game) FindFirstEmpty(c chan *Position) {
	go func() {
		defer close(c)

		for row := range g.Board {
			for col := range g.Board[0] {
				if g.Board[row][col] == 0 {
					c <- &Position{row, col}
				} else {
					continue
				}
			}
		}
		c <- nil
	}()
}


