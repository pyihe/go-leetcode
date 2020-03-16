package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
*/

func longestPalindrome(s string) string {
	if len(s) > 1000 {
		return ""
	}
	var maxLen, startIndex int
	for strLen := len(s); strLen >= 1; strLen-- {
		for i := 0; i+strLen-1 < len(s); i++ {
			tempStr := s[i : i+strLen]
			var j, k = 0, len(tempStr) - 1
			for j <= k {
				if tempStr[j] == tempStr[k] {
					j++
					k--
				} else {
					break
				}
				if j >= k {
					break
				}
			}
			if k <= j && maxLen <= strLen {
				maxLen = strLen
				startIndex = i
			}
		}
	}
	return s[startIndex : startIndex+maxLen]
}

func main() {
	fmt.Println(longestPalindrome("aacaa"))
}
