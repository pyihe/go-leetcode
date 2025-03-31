package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
   @Create by GoLand
   @Author: hong
   @Time: 2018/7/27 16:39
   @File: main.go
*/

/*
Question Description:
Given a 32-bit signed integer, reverse digits of an integer.
*/

func main() {
	fmt.Println(Reverse(1534236469))
}

func Reverse(x int) int {
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
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0
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
