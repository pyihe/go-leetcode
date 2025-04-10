package main

import (
	"fmt"
)

func canAliceWin(n int) bool {
	//step := 10
	//for n >= step {
	//	n, step = n-step, step-1
	//}
	//return step%2 != 0

	for step := 10; step > 0; step-- {
		n -= step
		if n < 0 {
			return step%2 != 0
		}
	}
	return false
}

func main() {
	fmt.Println(canAliceWin(12))
}
