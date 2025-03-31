package main

import "fmt"

/*
   @Create by GoLand
   @Author: hong
   @Time: 2018/7/24 18:03
   @File: main.go
*/

/*
Question Description:
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
*/

func main() {
	fmt.Println(longestPalindrome("abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa"))
}

//my solution
func longestPalindrome(s string) string {
	b := []byte(s)
	length := len(s)
	var maxLen int                 //the max len of palindromic substring
	var result string              //the result to return
	for i := 1; i <= length; i++ { //substring's len
		for j := 0; j+i <= length; j++ { //kind of substring for every substring len
			if isStrSymmetrical(string(b[j:j+i])) && i >= maxLen {
				maxLen = i
				result = string(b[j : j+i])
			}
		}
	}
	//the result should be []string under normal conditions
	return result
}

//judge the string s is or isn't palindromic string.
func isStrSymmetrical(s string) bool {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		if b[i] != b[len(b)-i-1] {
			return false
		}
	}
	return true
}

//leetcode go solution
//无
