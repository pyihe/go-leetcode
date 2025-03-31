package main

import "fmt"

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var newX int
	for i := x; i > 0; i = i / 10 {
		temp := i % 10
		newX = newX*10 + temp
	}
	return newX == x
}

func main() {
	fmt.Println(isPalindrome(121))
}
