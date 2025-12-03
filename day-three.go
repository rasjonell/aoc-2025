package main

import (
	"fmt"
	"strconv"

	"github.com/rasjonell/aoc/internal/input"
)

type BatterySum struct {
	total     uint64
	day, part uint8
}

func NewBatterySum(day, part uint8) ProblemSolution {
	return &BatterySum{day: day, part: part, total: 0}
}

func (b *BatterySum) Solve() string {
	banks, err := input.ReadLines(fmt.Sprintf("%d", b.day))
	if err != nil {
		return err.Error()
	}

	for _, bank := range banks {
		switch b.part {
		case 1:
			b.total += pickLargestK(bank, 2)
		case 2:
			b.total += pickLargestK(bank, 12)
		}
	}

	return fmt.Sprintf("%d", b.total)
}

func pickLargestTwo(numLine string) uint64 {
	type pick struct {
		val uint64
		idx int
	}
	first := pick{0, -1}
	second := pick{0, -1}

	for i, n := range numLine {
		num := fromRune(n)

		if i != len(numLine)-1 && (first.idx == -1 || num > first.val) {
			second = first
			first = pick{num, i}
			second = pick{fromRune(rune(numLine[i+1])), i + 1}
		} else if num == first.val {
			if second.idx == -1 || num > second.val {
				second = pick{num, i}
			}
		} else if second.idx == -1 || num > second.val {
			second = pick{num, i}
		} else {
		}
	}

	var stringRes string

	if first.idx < second.idx {
		stringRes = fmt.Sprintf("%d%d", first.val, second.val)
	} else {
		stringRes = fmt.Sprintf("%d%d", second.val, first.val)
	}

	return fromString(stringRes)
}

func pickLargestK(numLine string, k int) uint64 {
	toRemove := len(numLine) - k
	stack := make([]uint64, 0, k)

	for _, n := range numLine {
		digit := fromRune(n)

		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, digit)
	}

	str := ""
	for i := range k {
		str += fmt.Sprintf("%d", stack[i])
	}

	return fromString(str)
}

func fromRune(c rune) uint64 {
	num, err := strconv.ParseUint(fmt.Sprintf("%c", c), 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func fromString(s string) uint64 {
	result, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

