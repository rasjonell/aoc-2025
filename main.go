package main

import "fmt"

type ProblemSolution interface {
	Solve() string
}

func main() {
	// day 1 part 1
	// dial1 := NewDialSolution(100, 1, 1)
	// fmt.Println("Day 1 Part 1:", dial1.Solve())
	// day 1 part 2
	// dial2 := NewDialSolution(100, 1, 2)
	// fmt.Println("Day 1 Part 2:", dial2.Solve())

	// day 2 part 1
	validator1 := NewIDValidator(2, 1)
	fmt.Println("Day 2 Part 1", validator1.Solve())

	validator2 := NewIDValidator(2, 2)
	fmt.Println("Day 2 Part 2", validator2.Solve())
}
