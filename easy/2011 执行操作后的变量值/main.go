package main

import (
	"fmt"
)

func finalValueAfterOperations(operations []string) int {
	r := 0
	for _, op := range operations {
		switch op {
		case "X++", "++X":
			r += 1
		default:
			r -= 1
		}
	}
	return r
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"}))
}
