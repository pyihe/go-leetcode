package __TwoSum

/*
    @Create by GoLand
    @Author: hong
    @Time: 2018/7/19 17:51 
    @File: issue1.go    
*/

/*
Question Description:
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

//my solution
func TwoSum(nums []int, target int) []int {
	var result []int

	for index, element := range nums {
		for j := index + 1; j < len(nums); j++ {
			if element+nums[j] == target {
				result = append(result, index, j)
				break
			}
		}
	}
	return result
}

//go solution in leetcode
func TwoSum2(nums []int, target int) []int {
	var result []int
	numMap := make(map[int]int)

	for index, element := range nums {
		newTarget := target - element
		if newIndex, ok := numMap[newTarget]; ok {
			result = append(result, index, newIndex)
			break
		}
		numMap[element] = index
	}
	return result
}
