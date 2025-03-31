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
	//data := []int{-4, -1, -1, 0, 1, 2}
	//data := []int{0, 0, 0}
	//data := []int{-2, 0, 0, 2, 2}
	//fmt.Println(threeSum(data))

	data := []int{-3, 0, 1, 2}
	fmt.Println(threeSum(data))
}

//三数之和为0，则至少有一个数小于0。所以先对数组排序，挑选第一个数(第一个数必须小于0)，然后在从剩余对数组中筛选剩下的两个数，剩下两个的和加上第一个数==0则满足条件
//筛选剩余两个数时利用两个下标指针, 如果第一个数重复出现则跳过，因为如果该数符合条件，则会出现重复的答案
func threeSum(nums []int) (result [][]int) {
	if len(nums) < 3 {
		return
	}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v > 0 || i == len(nums)-2 {
			break
		}
		//这里防止结果里有重复的答案
		if i > 0 && nums[i-1] == v {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if v+nums[left]+nums[right] == 0 {
				result = append(result, []int{v, nums[left], nums[right]})
				for left+1 <= right && nums[left] == nums[left+1] {
					left++
				}
				for right-1 >= left && nums[right] == nums[right-1] {
					right--
				}
				right--
				left++
			} else if nums[left]+nums[right] > -v {
				right--
			} else {
				left++
			}
		}
	}
	return
}
