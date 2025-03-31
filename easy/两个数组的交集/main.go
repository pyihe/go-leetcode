package main

import "fmt"

func main() {
	arr1 := []int{4, 9, 5}
	arr2 := []int{9, 4, 9, 8, 4}
	fmt.Println(arrayBoth(arr1, arr2))

	//arr1 := []int{1,2,3,4,4,13}
	//arr2 := []int{1,2,3,9,10}
	//fmt.Println(arrayBothInOrder(arr1, arr2))
}

//这里有可能最先想到利用两个for循环，分别遍历两个数组中的元素，然后判断是否相等。这里存在的问题是：元素重复问题，

//349：求两个数组的公共元素，去重
func both(arr1, arr2 []int) []int {
	mp := make(map[int]bool)
	for i := range arr1 {
		mp[arr1[i]] = true
	}
	//这里复用arr2, 避免了一次内存申请
	k := 0
	for i := range arr2 {
		if mp[arr2[i]] {
			arr2[k] = arr2[i]
			delete(mp, arr2[i])
			k++
		}
	}
	return arr2[:k]
}

//利用map映射, 如果忽略元素的重复性，则map不需要记录次数，只需要存在即可
func arrayBoth(arr1, arr2 []int) []int {
	mp := make(map[int]int)
	for i := range arr1 {
		mp[arr1[i]] += 1
	}
	//这里复用arr2, 避免了一次内存申请
	k := 0
	for i := range arr2 {
		if mp[arr2[i]] > 0 {
			arr2[k] = arr2[i]
			mp[arr2[i]] -= 1
			k++
		}
	}
	return arr2[:k]
}

//如果两个数组有序：利用两个索引指针，如果相等则同时往后移动两个指针，否则只需要移动值较小的那个指针
func arrayBothInOrder(arr1, arr2 []int) []int {
	var index1, index2, k int
	for index1 < len(arr1) && index2 < len(arr2) {
		if arr1[index1] == arr2[index2] {
			arr1[k] = arr1[index1]
			index1++
			index2++
			k++
		} else if arr1[index1] > arr2[index2] {
			index2++
		} else {
			index1++
		}
	}
	return arr1[:k]
}
