package __MedianOfTwoSortedArrays

import (
	"fmt"
	"sort"
)

/*
   @Create by GoLand
   @Author: hong
   @Time: 2018/7/20 17:32
   @File: issue4.go
*/

/*
Question Description
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Sort(Ints(nums))
	return findMedian(nums)
}

func findMedian(nums []int) float64 {
	length := len(nums)
	if length%2 == 0 {
		return float64((nums[length/2] + nums[length/2-1])) / 2.0
	} else {
		return float64(nums[length/2])
	}
}

type Ints []int

func (s Ints) Len() int {
	return len(s)
}

func (s Ints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Ints) Less(i, j int) bool {
	return s[i] < s[j]
}

/*leetcode go solution*/
func findMedianSortedArrays2(a []int, b []int) float64 {
	m := len(a)
	n := len(b)

	l := m + n

	i, j, k := 0, 0, 0
	lastTwo := [2]int{}
	fmt.Println(i, j, k)
	for k <= l/2 && i < m && j < n {
		if a[i] < b[j] {
			lastTwo[0], lastTwo[1] = lastTwo[1], a[i]
			i++
		} else {
			lastTwo[0], lastTwo[1] = lastTwo[1], b[j]
			j++
		}
		k++
	}
	fmt.Println(i, j, k)
	for k <= l/2 && i < m {
		lastTwo[0], lastTwo[1] = lastTwo[1], a[i]
		i++
		k++
	}
	fmt.Println(i, j, k)
	for k <= l/2 && j < n {
		lastTwo[0], lastTwo[1] = lastTwo[1], b[j]
		k++
		j++
	}

	fmt.Println(i, j, k)
	if l%2 == 0 {
		return float64(lastTwo[0]+lastTwo[1]) / 2.0
	}

	return float64(lastTwo[1])
}
