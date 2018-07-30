package __ReverseInteger

import (
	"strconv"
	"math"
)

/*
    @Create by GoLand
    @Author: hong
    @Time: 2018/7/27 16:39 
    @File: issue7.go    
*/

/*
Question Description:
Given a 32-bit signed integer, reverse digits of an integer.
*/
func Reverse(x int) int {
	if x > math.MaxInt32 || x <= math.MinInt32 {
		return 0
	}
	xString := strconv.Itoa(x)
	var byteData []byte
	for i := len(xString) - 1; i >= 0; i-- {
		byteData = append(byteData, byte(xString[i]))
	}
	if string(byteData[len(byteData)-1]) == "-" {
		byteData = byteData[:len(byteData)-1]
		byteData = append([]byte("-"), byteData...)
	}
	n, err := strconv.Atoi(string(byteData))
	if err != nil {
		panic(err)
	}
	return n
}

//leetcode go solution
func reverse(x int) int {
	var resp int

	for x != 0 {
		resp = resp*10 + x%10
		x /= 10
		if resp > math.MaxInt32 || resp < math.MinInt32 {
			return 0
		}
	}

	return resp
}
