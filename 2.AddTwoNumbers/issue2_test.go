package __AddTwoNumbers

import (
	"fmt"
	"testing"
)

/*
   @Create by GoLand
   @Author: hong
   @Time: 2018/7/19 18:11
   @File: issue2_test.go
*/

func TestGetNumber(t *testing.T) {
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  4,
		Next: l3,
	}

	l1 := &ListNode{
		Val:  2,
		Next: l2,
	}

	l4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  6,
		Next: l4,
	}
	l6 := &ListNode{
		Val:  5,
		Next: l5,
	}
	l7 := &ListNode{
		Val:  9,
		Next: l6,
	}

	result := addTwoNumbers(l1, l7)
	fmt.Println(fmt.Sprintf("result = %+v", *result))

label:
	if result.Next != nil {
		result = result.Next
		fmt.Println(fmt.Sprintf("%+v", *result))
		goto label
	}

}
