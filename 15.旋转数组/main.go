package main

import "fmt"

func main() {
	var data = []int{1,2,3,4,5,6,7}
	rotate(data, 3)
	fmt.Println(data)
}


//三次翻转
func rotate(nums []int, k int)  {
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
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}