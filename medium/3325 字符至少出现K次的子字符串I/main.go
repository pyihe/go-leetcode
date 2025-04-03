package main

import (
	"fmt"
)

func numberOfSubstrings(s string, k int) int {
	var (
		left   = 0
		cnt    = 0
		result = [26]int{}
	)
	for _, r := range s {
		result[r-'a']++
		for result[r-'a'] >= k {
			result[s[left]-'a']--
			left++
		}
		cnt += left
	}
	return cnt
}

func main() {
	fmt.Println(numberOfSubstrings("ajsrhoebe", 2))
}
