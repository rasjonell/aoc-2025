package main

import "fmt"

type ProblemSolution interface {
	Solve(inputFile string) string
}

func main() {
	// day 1 part 1
	dial1 := NewDialSolution(100, 1)
	fmt.Println("Day 1 Part 1:", dial1.Solve("1"))
	dial2 := NewDialSolution(100, 2)
	fmt.Println("Day 1 Part 2:", dial2.Solve("1"))
}
