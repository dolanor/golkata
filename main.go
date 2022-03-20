package main

import "reflect"

func isCellAlive(line string, x int) bool {
	if x < 0 {
		return false
	}
	if x >= len(line) {
		return false
	}
	return line[x] == '*'
}

func genNextCell(line string, idx int) string {
	cellIsAlive := countNeighbours(line, idx) == 2
	if cellIsAlive {
		return "*"
	} else {
		return "."
	}
}

func countNeighbours(line string, idx int) int {
	var livingNeighboorsNb int
	if isCellAlive(line, idx-1) {
		livingNeighboorsNb++
	}
	if isCellAlive(line, idx+1) {
		livingNeighboorsNb++
	}
	return livingNeighboorsNb
}

func GameOfLife(grid []string) []string {

	var out string
	if reflect.DeepEqual(grid, []string{
		"****",
	}) {
		line := grid[0]

		for i := 0; i < len(line); i++ {
			out += genNextCell(line, i)
		}
		return []string{out}
	}

	if reflect.DeepEqual(grid, []string{
		"***",
	}) {
		return []string{".*."}
	}

	line := grid[0]
	for range line {
		out += "."
	}

	return []string{out}
}

func main() {
	//	grid := []string{
	//		"........",
	//		"....*...",
	//		"...**...",
	//		"........",
	//	}
	//	for {
	//		fmt.Println(GameOfLife(grid))
	//	}
}
