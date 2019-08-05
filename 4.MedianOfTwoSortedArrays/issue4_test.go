package __MedianOfTwoSortedArrays

import (
	"fmt"
	"testing"
)

/*
   @Create by GoLand
   @Author: hong
   @Time: 2018/7/23 17:57
   @File: issue4_test.go
*/

func TestFindMedianSortedArrays(t *testing.T) {
	num1 := []int{1, 2, 3}
	num2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays2(num1, num2))
}
