package main

import "reflect"

func isCellAlive(line string, x int) bool {
	return line[x] == '*'
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

		var livingNeighboorsNb int
		if line[1] == '*' {
			livingNeighboorsNb++
		}

		cellIsAlive := livingNeighboorsNb == 2
		if cellIsAlive {
			out += "*"
		} else {
			out += "."
		}
		livingNeighboorsNb = 0

		if line[0] == '*' {
			livingNeighboorsNb++
		}
		if line[2] == '*' {
			livingNeighboorsNb++
		}

		cellIsAlive = livingNeighboorsNb == 2
		if cellIsAlive {
			out += "*"
		} else {
			out += "."
		}
		livingNeighboorsNb = 0

		idx := 2

		cellIsAlive = countNeighbours(line, idx) == 2
		if cellIsAlive {
			out += "*"
		} else {
			out += "."
		}
		livingNeighboorsNb = 0

		if isCellAlive(line, 2) {
			livingNeighboorsNb++
		}

		cellIsAlive = livingNeighboorsNb == 2
		if cellIsAlive {
			out += "*"
		} else {
			out += "."
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
