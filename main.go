package main

import (
	"fmt"
	p "n-puzzle/parsing"
	s "n-puzzle/solver"
)

func main() {

	size, heuristic := p.CheckFlags()
	var Puzzle []int
	if size == 0 {
		Puzzle, size = p.ReadBoardFromFile(Puzzle, size)
	} else {
		Puzzle = p.GenerateRandomBoard(size)
	}
	s.Solver(Puzzle, size, heuristic)

	fmt.Println("\nYou've finished n-puzzle!")

}
