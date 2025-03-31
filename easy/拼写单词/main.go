package main

import "strings"

/*
给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。

假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。

注意：每次拼写时，chars 中的每个字母都只能用一次。

返回词汇表 words 中你掌握的所有单词的 长度之和。
*/

func countCharacters(words []string, chars string) int {
	var result int
	var ok = true
	for _, w := range words {
		for _, rw := range w {
			temp := string(rw)
			if strings.Count(chars, temp) < strings.Count(w, temp) {
				ok = false
				break
			}
		}
		if ok {
			result += len(w)
		}
		ok = true
	}
	return result

	/*
	var cMap = make(map[rune]int)
		var wMap = make(map[rune]int)
		for _, cr := range chars {
			cMap[cr] += 1
		}
		var result int
		var ok = true
		for _, w := range words {
			for _, rw := range w {
				wMap[rw] += 1
			}
			for k := range wMap {
				if cMap[k] < wMap[k] {
					ok = false
				}
			}
			if ok {
				result += len(w)
			}
			ok = true
			wMap = make(map[rune]int)
		}
		return result
	*/
}
