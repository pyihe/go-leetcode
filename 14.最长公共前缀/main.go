package main

import (
	"fmt"
	"strings"
)

func main() {
	src := []string{"c", "acc", "ccc"}
	//src := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(src))
}

//不通过的solution, 遍历每个元素的每个位置，如果有不同的则退出，否则继续下一趟的比较
func longestPrefix(src []string) string {
	var result  []byte
	 k := 0

	for {
		for j := range src {
			if len(src[j]) == k {
				return string(result[:k])
			}
			if j == 0 {
				result = append(result, src[j][k])
			}
			if src[j][k] != result[k] {
				return string(result[:k])
			}
		}
		k++
	}
}

//假设第一个元素就是最长公共前缀，则用第一个元素与其他比较，如果不同，则去掉公共最长前缀的最后一个元素后再次比较
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs {
		for len(prefix) > 0 {
			if strings.Index(str, prefix) == 0 {
				break
			}
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			return prefix
		}
	}
	return prefix
}