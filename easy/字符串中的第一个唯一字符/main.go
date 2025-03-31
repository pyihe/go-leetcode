package main

import (
	"fmt"
	"strings"
)

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。


示例：
s = "leetcode"
返回 0

s = "loveleetcode"
返回 2


提示：你可以假定该字符串只包含小写字母。
*/

func main() {
	fmt.Printf("%v\n", firstUniqChar("dddccdbba"))
}

//my solution
func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}
	var m = make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if !strings.ContainsAny(s[i+1:], string(s[i])) && !m[s[i]] {
			return i
		}
		m[s[i]] = true
	}
	return -1
}

func solution2(s string) int {
	var m [26]int
	for i, v := range s {
		m[v - 'a'] = i
	}
	for i, v := range s {
		if m[v - 'a'] == i {
			return i
		} else {
			m[v-'a'] = -1
		}
	}
	return -1
}