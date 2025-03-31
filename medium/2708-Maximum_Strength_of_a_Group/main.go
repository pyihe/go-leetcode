package main

import (
	"fmt"
	"slices"
)

func maxStrength(nums []int) int64 {
	var (
		zero    int
		pos     int
		r       = 1
		negList = make([]int, 0, len(nums))
	)
	for _, n := range nums {
		switch {
		case n == 0:
			zero += 1
		case n > 0:
			pos += 1
		default:
			negList = append(negList, n)
		}
		if n != 0 {
			r *= n
		}
	}
	if len(nums) == 1 {
		return int64(nums[0])
	}
	if pos == 0 && len(negList) <= 1 {
		return 0
	}
	if r > 0 {
		return int64(r)
	}
	slices.SortFunc(negList, func(a, b int) int {
		if a < b {
			return 1
		}
		if a > b {
			return -1
		}
		return 0
	})
	return int64(r / negList[0])
}

func main() {
	fmt.Println(maxStrength([]int{-4, -5, -4}))
}
