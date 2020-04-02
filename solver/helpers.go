package solver

import "math"

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
