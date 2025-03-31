package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func main() {
	data := []int{-2, -1, 0, 0, 1, 2}
	fmt.Println(fourSum(data, 0))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var result [][]int
	if len(nums) < 4 {
		return result
	}

	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		if i == len(nums)-3 {
			return result
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j-1] == nums[j] {
				continue
			}
			n1 := nums[i]
			n2 := nums[j]
			left := j + 1
			right := len(nums) - 1
			for left < right {
				n3 := nums[left]
				n4 := nums[right]
				sum := n1 + n2 + n3 + n4
				if sum == target {
					result = append(result, []int{n1, n2, n3, n4})
					//排除重复的答案
					for left+1 <= right && nums[left] == nums[left+1] {
						left++
					}
					for right-1 >= left && nums[right] == nums[right-1] {
						right--
					}
					right--
					left++
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}
