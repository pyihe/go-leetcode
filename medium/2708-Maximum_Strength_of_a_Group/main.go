package main

import (
	"fmt"
	"math"
)

func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	var (
		zeros  int
		pos    int
		negs   int
		maxNeg = math.MinInt
		r      = 1
	)
	for _, n := range nums {
		switch {
		case n == 0:
			zeros += 1
		case n > 0:
			pos += 1
		default:
			negs += 1
			if n > maxNeg {
				maxNeg = n
			}
		}
		if n != 0 {
			r *= n
		}
	}
	if pos == 0 && negs <= 1 {
		return 0
	}
	if r > 0 {
		return int64(r)
	}
	return int64(r / maxNeg)
}

func main() {
	fmt.Println(maxStrength([]int{-4, -5, -4}))
}
