package main

import "fmt"

/*
   @Create by GoLand
   @Author: hong
   @Time: 2018/7/19 17:51
   @File: main.go
*/

/*
Question Description:
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

func main() {
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


//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//my solution
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var listArray []*ListNode
	flag := false
label:
	l1Val := -1
	l2Val := -1
	tempNode := &ListNode{}
	if l1 != nil {
		l1Val = l1.Val
	}
	if l2 != nil {
		l2Val = l2.Val
	}

	if flag && l1Val == -1 && l2Val == -1 {
		tempNode.Val = 1
		listArray = append(listArray, tempNode)
	}
	if l1Val != -1 && l2Val != -1 {
		num := l1Val + l2Val
		if !flag {
			tempNode.Val = (num) % 10
			flag = num >= 10
		} else {
			tempNode.Val = (num + 1) % 10
			flag = (num + 1) >= 10
		}
		listArray = append(listArray, tempNode)
		l1 = l1.Next
		l2 = l2.Next
		goto label
	}
	if l1Val == -1 && l2Val != -1 {
		if flag {
			tempNode.Val = (l2Val + 1) % 10
			flag = (l2Val + 1) >= 10
		} else {
			tempNode.Val = l2Val % 10
		}
		listArray = append(listArray, tempNode)
		//l1 = l1.Next
		l2 = l2.Next
		goto label
	}
	if l1Val != -1 && l2Val == -1 {
		if flag {
			tempNode.Val = (l1Val + 1) % 10
			flag = (l1Val + 1) >= 10
		} else {
			tempNode.Val = l1Val
		}
		listArray = append(listArray, tempNode)
		l1 = l1.Next
		//l2 = l2.Next
		goto label
	}

	for i := 0; i < len(listArray)-1; i++ {
		listArray[i].Next = listArray[i+1]
	}
	return listArray[0]
}

//leetcode go solution: 递归
func helper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	var l1N, l2N *ListNode
	l1V, l2V := 0, 0

	// Termination condition
	if l1 == nil && l2 == nil {
		if carry != 0 {
			return &ListNode{carry, nil}
		}
		return nil
	}

	if l1 != nil {
		l1N = l1.Next
		l1V = l1.Val
	}

	if l2 != nil {
		l2N = l2.Next
		l2V = l2.Val
	}

	sum := l1V + l2V + carry
	n := &ListNode{sum % 10, nil}
	n.Next = helper(l1N, l2N, sum/10)
	return n
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	return helper(l1, l2, 0)
}
