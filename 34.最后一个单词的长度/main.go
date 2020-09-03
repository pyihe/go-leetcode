package main

import (
	"fmt"
	"strings"
)

/*
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。

示例:

输入: "Hello World"
输出: 5
*/

func main() {
	fmt.Println(lengthOfLastWord1("Hello World"))
}

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	n := len(arr) - 1
	for n >= 0 {
		if len(arr[n]) > 0 {
			return len(arr[n])
		}
		n--
	}
	return 0
}

func lengthOfLastWord1(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		if i == 0 {
			return 1
		}
		for k := i - 1; k >= 0; k-- {
			if s[k] == ' ' {
				return len(s[k+1 : i+1])
			}
		}
		return len(s[:i+1])
	}
	return 0
}
