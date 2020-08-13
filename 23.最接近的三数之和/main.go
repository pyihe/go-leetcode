package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 

提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4
*/

func main() {
	data := []int{-3, 0, 1, 2}
	fmt.Println(threeSumClosest(data, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var interval = math.MaxFloat64
	var result int
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if i == len(nums)-2 {
			break
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := v + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			temp := math.Abs(float64(sum - target))
			if temp < interval {
				interval = temp
				result = sum
			}
			if sum < target {
				for left+1 <= right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				for right-1 >= left && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
	return result
}
