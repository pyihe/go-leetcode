package main

import (
	"fmt"
)

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。
*/

func main() {
	var data = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(data, 3)
	fmt.Println(data)
}

//三次翻转
func rotate(nums []int, k int) {
	if k <= 0 {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}

func reverse(arr []int) {
	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {
		arr[left], arr[right] = arr[right], arr[left]
	}
}

func rotate1(nums []int, k int) {
	var pre, temp int
	for i := 0; i < k; i++ {
		pre = nums[len(nums) -1]
		for j := 0; j < len(nums); j++ {
			temp = nums[j]
			nums[j] = pre
			pre = temp
		}
	}
}