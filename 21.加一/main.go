package main

import "fmt"

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/

func main() {
	data := []int{2, 4, 9, 3, 9}
	fmt.Println(plusOne(data, 1, 10))
}

//addNum: 增量, scale: 进制
func plusOne(digits []int, addNum int, scale int) []int {
	if scale <= 0 || addNum >= scale {
		return digits
	}
	//不需要进位
	if digits[len(digits)-1]+addNum <= scale-1 {
		digits[len(digits)-1] += addNum
		return digits
	}
	//需要进位
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+addNum >= scale {
			digits[i] = (digits[i] + addNum) % scale
			addNum = 1
			if i == 0 { //需要拓展数组长度
				digits = append([]int{addNum}, digits...)
			}
		} else {
			digits[i] = digits[i] + addNum
			break
		}
	}
	return digits
}
