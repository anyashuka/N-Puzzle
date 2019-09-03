package solver

import (
	"bytes"
	"container/heap"
	"fmt"
	"log"
	g "n-puzzle/golib"
	"os"
	"time"

	"github.com/AndreasBriese/bbloom"
	// "time"
)

type Problem struct {
	start []int
	goal  []int
<<<<<<< HEAD
=======
	// solvable bool
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
	//heuristic      string
	//searchAlgo     string
	solutionPath   map[int][]int // maybe unnecessary?
	sizeComplexity int
	timeComplexity int
	solutionFound  bool
}

func newProblem(Puzzle []int, size int) Problem {
	problem := Problem{}
	problem.start = Puzzle
	problem.goal = MakeGoal(size)
<<<<<<< HEAD
=======
	// problem.solvable = IsSolvable(problem.goal, Puzzle, size)
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
	//problem.heuristic = "MANHATTAN"
	//problem.searchAlgo = "A_STAR"
	problem.sizeComplexity = 0
	problem.timeComplexity = 1
	return problem
}

type State struct {
	index     int
	priority  int
	depth     int
	heuristic int
	puzzle    []int
}

func newState(Puzzle []int, priority int, depth int, heuristic int) *State {
	state := &State{}
	state.index = 0
	state.priority = priority
<<<<<<< HEAD
	state.depth = depth         // not sure if we need to store this?
	state.heuristic = heuristic // not sure about this one either?
=======
	state.depth = depth         // not sure if I need to store this either?
	state.heuristic = heuristic // I don't think I need to keep this?
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
	state.puzzle = Puzzle
	return state
}

<<<<<<< HEAD
func Solver(Puzzle []int, size int) {
	// TESTING RUNTIME
	start := time.Now()

	problem := newProblem(Puzzle, size)
	goal := g.PuzzleToString(problem.goal, ",")
=======
func Solver(Puzzle []int, size int, iterations int) {
	start := time.Now()

	problem := newProblem(Puzzle, size)
	unsolved := true

	goal := g.PuzzleToString(problem.goal, ",")
	// g.PrintBoard(Puzzle, size)
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c

	if IsSolvable(problem.goal, Puzzle, size) == false {
		fmt.Println("This puzzle in unsolvable.")
		os.Exit(1)
	}

	state := newState(Puzzle, 100000, 0, 0)

	openSet := make(map[string]int)
	parent := g.PuzzleToString(state.puzzle, ",")
	openSet[parent] = state.priority

<<<<<<< HEAD
=======
	// closedSet := make(map[string]int)
	closedSet := bbloom.New(float64(1<<16), float64(0.01))
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
	openQueue := CreateQueue(*state)
	closedSet := bbloom.New(float64(1<<16), float64(0.01))

<<<<<<< HEAD
	unsolved := true
=======
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
	for unsolved {

		if len(openQueue) == 0 {
			fmt.Println("This priorityQueue is empty.")
			g.PrintBoard(state.puzzle, size)
			os.Exit(1)
		}

		state = heap.Pop(&openQueue).(*State)
		parent = g.PuzzleToString(state.puzzle, ",")
		delete(openSet, parent)

		closedSet.AddIfNotHas([]byte(parent))
<<<<<<< HEAD
=======
		// if _, exists := closedSet[parent]; !exists {
		// 	closedSet[parent] = state.priority
		// }
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c

		if bytes.Equal([]byte(parent), []byte(goal)) {
			fmt.Println("This puzzle has been solved!\n")
			g.PrintBoard(state.puzzle, size)
			// REBUILD PATH TO START
<<<<<<< HEAD

			// TESTING RUNTIME
			elapsed := time.Since(start)
			log.Printf("Binomial took %s", elapsed)
			unsolved = false
			// os.Exit(1)
=======
			elapsed := time.Since(start)
			log.Printf("Binomial took %s", elapsed)
			os.Exit(1)
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
		}

		// if reflect.DeepEqual(problem.goal, state.puzzle) {
		// fmt.Println("This puzzle has been solved!\n")
		// g.PrintBoard(state.puzzle, size)
		// REBUILD PATH TO START
		// unsolved = false
		// elapsed := time.Since(start)
		// log.Printf("Binomial took %s", elapsed)
		// os.Exit(1)
		// }
<<<<<<< HEAD

		children := CreateNeighbors(state.puzzle, size)
=======

		// fmt.Printf("\n-- parent --")
		// g.PrintBoard(state.puzzle, size)
		// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d\n", state.priority, state.heuristic, state.depth)

		// time.Sleep(1000 * time.Millisecond)

		children := CreateNeighbors(state.puzzle, size)
		// fmt.Printf("\n-- child --")
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c

		for _, child := range children {
			tmpChild := g.PuzzleToString(child, ",")

			if bytes.Equal([]byte(parent), []byte(tmpChild)) {
				fmt.Println("This puzzle has been solved!\n")
				g.PrintBoard(state.puzzle, size)
				// REBUILD PATH TO START

				// TESTING RUNTIME
				elapsed := time.Since(start)
				log.Printf("Binomial took %s", elapsed)
				unsolved = false
			}

			// if reflect.DeepEqual(problem.goal, child) {
			// fmt.Println("This puzzle has been solved!\n")
			// g.PrintBoard(child, size)
<<<<<<< HEAD
=======
			tmpChild := g.PuzzleToString(child, ",")

			if bytes.Equal([]byte(parent), []byte(tmpChild)) {
				fmt.Println("This puzzle has been solved!\n")
				g.PrintBoard(state.puzzle, size)
				// REBUILD PATH TO START
				elapsed := time.Since(start)
				log.Printf("Binomial took %s", elapsed)
				os.Exit(1)
			}

			// if reflect.DeepEqual(problem.goal, child) {
			// fmt.Println("This puzzle has been solved!\n")
			// g.PrintBoard(child, size)
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
			// REBUILD PATH TO START
			// elapsed := time.Since(start)
			// log.Printf("Binomial took %s", elapsed)
			// os.Exit(1)
			// }

			if closedSet.Has([]byte(tmpChild)) {
				problem.sizeComplexity++
				continue
			}

<<<<<<< HEAD
			depth := -(state.depth + 1)
			// depth = -depth
			heuristic := g.Manhattan(child, problem.goal, size)
			s := newState(child, depth+heuristic, depth, heuristic)

=======
			depth := state.depth + 1
			depth = -depth
			heuristic := g.Manhattan(child, problem.goal, size)
			// priority = -priority
			// priority :=  heuristic
			s := newState(child, depth+heuristic, depth, heuristic)

			// if _, exists := closedSet[tmpChild]; exists {
			// 	continue
			// }
			// g.PrintBoard(child, size)
			// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d\n", s.priority, s.heuristic, s.depth)

>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
			if _, exists := openSet[tmpChild]; exists {
				if openSet[tmpChild] < s.priority {
					continue
				}
			}

			openSet[tmpChild] = s.priority
			heap.Push(&openQueue, s)

			// problem.timeComplexity++

<<<<<<< HEAD
=======
			// s := newState(child, size, priority, state.depth+1, heuristic)

>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
		}
	}
}
<<<<<<< HEAD
=======

// if bytes.Equal([]byte(g.PuzzleToString(problem.goal, ",")), []byte(g.PuzzleToString(state.puzzle, ","))) {
// 	fmt.Println("This puzzle has been solved!\n")
// 	g.PrintBoard(state.puzzle, size)
// 	// REBUILD PATH TO START
// 	os.Exit(1)
// }

// fmt.Println(" ----------- ")
// fmt.Println("\n NEW STATE")
// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d", state.priority, state.heuristic, state.depth)
// g.PrintBoard(state.puzzle, size)
// fmt.Println(" ----------- ")

// if s.priority > state.priority {
// 	continue
// }
// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d", priority, heuristic, state.depth + 1)
// g.PrintBoard(child, size)
>>>>>>> d35897d97dcb5c31ed711cfb14b0856fdbba3d0c
