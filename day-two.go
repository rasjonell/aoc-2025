package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rasjonell/aoc/internal/input"
)

type IDValidator struct {
	total     uint64
	day, part uint8
}

func NewIDValidator(day, part uint8) ProblemSolution {
	return &IDValidator{day: day, part: part, total: 0}
}

func (v *IDValidator) Solve() string {
	var err error
	input.ReadSeperator(fmt.Sprintf("%d", v.day), ',', func(numRange string) {
		bounds := strings.Split(numRange, "-")
		l, lerr := strconv.ParseUint(bounds[0], 10, 64)
		if lerr != nil {
			err = lerr
			return
		}
		r, rerr := strconv.ParseUint(bounds[1], 10, 64)
		if rerr != nil {
			err = rerr
			return
		}
		for n := l; n <= r; n++ {
			switch v.part {
			case 1:
				if isDoubleRepeat(n) {
					v.total += n
				}
			case 2:
				if isMultiRepeat(n) {
					v.total += n
				}
			}
		}
	})

	if err != nil {
		return err.Error()
	}

	return strconv.FormatUint(v.total, 10)
}

func isDoubleRepeat(n uint64) bool {
	s := strconv.FormatUint(n, 10)
	if len(s)%2 != 0 {
		return false
	}
	h := len(s) / 2
	return s[:h] == s[h:]
}

func isMultiRepeat(n uint64) bool {
	s := strconv.FormatUint(n, 10)
	if len(s) < 2 {
		return false
	}

	double := s + s
	return strings.Contains(double[1:len(double)-1], s)
}
