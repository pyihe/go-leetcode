package main

/*
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

 

示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
*/

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//利用两个竞速指针(快慢指针)遍历，如果存在环，则快指针必定会与慢指针相遇
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1 := head.Next
	for head != nil && p1 != nil && p1.Next != nil{
		if head == p1 {
			return true
		}
		head = head.Next
		p1 = p1.Next.Next
	}
	return false
}
