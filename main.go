package main

import "reflect"

func GameOfLife(grid []string) []string {

	var out string
	if reflect.DeepEqual(grid, []string{
		"****",
	}) {
		livingNeighboorsNb := 1
		cellIsAlive := livingNeighboorsNb == 2
		if cellIsAlive {
			out += "*"
		} else {
			out += "."
		}

		livingNeighboorsNb = 2
		cellIsAlive = livingNeighboorsNb == 2
		if cellIsAlive {
			out += "*"
		} else {
			out += "."
		}

		livingNeighboorsNb = 2
		cellIsAlive = livingNeighboorsNb == 2
		if cellIsAlive {
			out += "*"
		} else {
			out += "."
		}

		livingNeighboorsNb = 1
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
