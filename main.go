package main

import "fmt"

type ProblemSolution interface {
	Solve(inputFile string) string
}

func main() {
	// day 1 part 1
	dial := NewDialSolution(100)
	fmt.Println("Day 1 Part 1:", dial.Solve("1"))
}
