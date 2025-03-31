package main

func percentageLetter(s string, letter byte) int {
	n := len(s)
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == letter {
			cnt++
		}
	}
	return 100 * cnt / n
}
