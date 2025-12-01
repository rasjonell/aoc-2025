package main

import (
	"strconv"

	"github.com/rasjonell/aoc/internal/input"
)

type Dial struct {
	fullRotate uint
	pos        uint64
	n          uint64
}

type Direction = rune

func NewDialSolution(n uint64) ProblemSolution {
	return &Dial{n: n, pos: 50, fullRotate: 0}
}

func (d *Dial) Process(line string) *Dial {
	dir := Direction(line[0])
	amountStr, err := strconv.ParseUint(line[1:], 10, 64)
	if err != nil {
		panic(err)
	}

	d.rotate(dir, amountStr)

	return d
}

func (d *Dial) rotate(dir Direction, amount uint64) {
	amount %= d.n
	if dir == 'R' {
		d.pos = (d.pos + amount) % d.n
	} else {
		d.pos = (d.pos + (d.n - amount)) % d.n
	}

	if d.pos == 0 {
		d.fullRotate++
	}
}

func (d *Dial) Solve(inputFile string) string {
	operations, err := input.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	for _, op := range operations {
		d.Process(op)
	}

	return strconv.FormatUint(uint64(d.fullRotate), 10)
}
