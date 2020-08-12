package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func main() {
	data := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(data))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		left := i+1
		right := len(nums)-1
		if nums[i] > 0 {
			continue
		}
		if i == 0 || nums[i] != nums[i-1] {
			for left < right {
				if nums[left] + nums[right] == target {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right++
				} else if nums[left] + nums[right] < target {
					left++
				} else {
					right++
				}
			}
		}
	}
	return result
}
//TODO |22|[15](https://leetcode-cn.com/problems/3sum/)|[三数之和]()|
