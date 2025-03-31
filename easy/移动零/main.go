package main

import "fmt"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/

func main() {
	var src = []int{0, 1, 0, 3, 12}
	moveZeroes(src)
	fmt.Println(src)
}

//思路：将非零元素挨个移动到数组前面，最后从非零元素下一个开始到最后一个元素，统统值为零
func moveZeroes(nums []int) {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}
