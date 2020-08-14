package main

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//两个指针，指针之间间隔n-1个节点，这样其中一个节点遍历到最后时，另一个节点则正好处于需要删除的节点。同时在遍历过程中用pre指向当前节点的前一个节点，用于删除操作
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//指向头节点的哨兵节点，用于保存头节点
	var guard = &ListNode{
		Next: head,
	}
	var pre *ListNode
	var current = guard //current最终指向需要删除的节点
	counter := 1
	for head != nil {
		//当遍历到第n个节点时，开始移动current指针，保证current和head之间差n-1个节点
		if counter >= n {
			pre = current
			current = current.Next
		}
		//head一直向后移动直到最后一个节点
		head = head.Next
		counter++
	}
	pre.Next = pre.Next.Next
	return guard.Next
}

//正向删除：利用两次遍历，第一次获取节点总数，第二次执行删除
func remove2(head *ListNode, n int) *ListNode {
	var guard = &ListNode{
		Next: head,
	}
	total := 1 //节点总数
	for head != nil {
		total++
		head = head.Next
	}

	//删除倒数第n个节点，等价于正向删除第total-n+1个节点
	head = guard.Next
	var counter = 1
	var pre = guard
	for head != nil {
		if counter == total-n+1 {
			//执行删除
			pre.Next = pre.Next.Next
			break
		}
		pre = pre.Next
		head = head.Next
		counter++
	}
	return guard.Next
}
