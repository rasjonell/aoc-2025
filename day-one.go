package main

import (
	"strconv"

	"github.com/rasjonell/aoc/internal/input"
)

type Dial struct {
	fullRotate uint64
	pos        uint64
	n          uint64
	part       uint8
}

type Direction = rune

func NewDialSolution(n uint64, part uint8) ProblemSolution {
	return &Dial{n: n, pos: 50, fullRotate: 0, part: part}
}

func (d *Dial) process(line string) *Dial {
	dir := Direction(line[0])
	amount, err := strconv.ParseUint(line[1:], 10, 64)
	if err != nil {
		panic(err)
	}

	switch d.part {
	case 1:
		d.rotatePartOne(dir, amount)
	case 2:
		d.rotatePartTwo(dir, amount)
	}

	return d
}

func (d *Dial) rotatePartOne(dir Direction, amount uint64) {
	amount %= d.n
	switch dir {
	case 'R':
		d.pos = (d.pos + amount) % d.n
	case 'L':
		d.pos = (d.pos + (d.n - amount)) % d.n
	}

	if d.pos == 0 {
		d.fullRotate++
	}
}

func (d *Dial) rotatePartTwo(dir Direction, amount uint64) {
	full := amount
	step := full % d.n
	var rotations uint64

	switch dir {
	case 'R':
		rotations = (d.pos + full) / d.n
		d.pos = (d.pos + step) % d.n
	case 'L':
		dist := (d.n - d.pos) % d.n
		rotations = (dist + full) / d.n
		d.pos = (d.pos + d.n - step) % d.n
	}

	d.fullRotate += rotations
}

func (d *Dial) Solve(inputFile string) string {
	operations, err := input.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	for _, op := range operations {
		d.process(op)
	}

	return strconv.FormatUint(uint64(d.fullRotate), 10)
}
