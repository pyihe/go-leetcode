package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
   @Create by GoLand
   @Author: hong
   @Time: 2018/8/7 17:20
   @File: main.go
*/

/*
Question Description:
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
*/

func main() {
	fmt.Println(myAtoi("+1"))
}

func myAtoi(str string) int {
	var bStr []byte
	for _, s := range str {
		if s == 32 {
			if len(bStr) > 0 {
				break
			}
			continue
		}
		if s == 45 { //-
			if len(bStr) > 0 {
				break
			}
			bStr = append(bStr, byte(s))
			continue
		}
		if s == 43 {
			if len(bStr) > 0 {
				break
			}
			bStr = append(bStr, byte(s))
			continue
		}
		if s < 48 || s > 57 {
			break
		}
		bStr = append(bStr, byte(s))
	}
	fmt.Println(string(bStr))
	switch len(bStr) {
	case 0:
		return 0
	case 1:
		if bStr[0] == byte('-') || bStr[0] == byte('+') {
			return 0
		} else {
			n, _ := strconv.Atoi(string(bStr))
			return n
		}
	default:
		if bStr[0] == byte('-') && bStr[1] == byte('+') {
			return 0
		}
		if bStr[1] == byte('-') && bStr[0] == byte('+') {
			return 0
		}
		if bStr[0] == byte('-') {
			n, _ := strconv.Atoi(string(bStr[1:]))
			n = 0 - n
			if n < math.MinInt32 {
				return math.MinInt32
			}
			if n > math.MaxInt32 {
				return math.MaxInt32
			}
			return n
		} else if bStr[0] == byte('+') {
			n, _ := strconv.Atoi(string(bStr[1:]))
			if n < math.MinInt32 {
				return math.MinInt32
			}
			if n > math.MaxInt32 {
				return math.MaxInt32
			}
			return n
		} else {
			n, _ := strconv.Atoi(string(bStr))
			if n < math.MinInt32 {
				return math.MinInt32
			}
			if n > math.MaxInt32 {
				return math.MaxInt32
			}
			return n
		}

	}
	return 0
}

/*leetcode go solution*/
func MyAtoi(str string) int {
	var i, num int
	sign := 1
	if str == "" {
		return 0
	}
	for i < len(str) && str[i] == 32 {
		i++
	}
	if i >= len(str) {
		return 0
	}
	if str[i] == 43 {
		i++
	} else if str[i] == 45 {
		sign = -1
		i++
	}
	for ; i < len(str); i++ {
		if str[i] != 0 && (str[i] < 48 || str[i] > 57) {
			return num * sign
		}
		n, _ := strconv.Atoi(string(str[i]))
		num = num*10 + n
		if num*sign < math.MinInt32 {
			return math.MinInt32
		} else if num*sign > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return num * sign
}
