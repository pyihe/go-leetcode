package main

import "fmt"

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

func main() {
	fmt.Printf("%d\n", 7/10)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//思路：设置哨兵节点，并添加一个游标指针，从l1, l2的头节点开始比较，将较小的那一个放进目标链表中
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	guard := &ListNode{}
	head := guard

	for l1 != nil || l2 != nil {
		if l1 == nil {
			head.Next = l2
			l2 = l2.Next
			head = head.Next
			continue
		}
		if l2 == nil {
			head.Next = l1
			l1 = l1.Next
			head = head.Next
			continue
		}
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	return guard.Next
}
