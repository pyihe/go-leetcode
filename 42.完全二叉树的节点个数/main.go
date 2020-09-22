package main

import (
	"container/list"
	"fmt"
)

/*
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~2h个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6
*/

func main() {
	var root = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	var num = []int{2, 3, 4, 5, 6}
	for _, v := range num {
		insert(root, v)
	}
	fmt.Printf("%v\n",countNodes(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(root *TreeNode, key int) {
	if root == nil {
		return
	}
	switch {
	case root.Val-key > 0: //插入做分支
		if root.Left == nil {
			var node = &TreeNode{
				Val:   key,
				Left:  nil,
				Right: nil,
			}
			root.Left = node
			return
		}
		insert(root.Left, key)
	case root.Val-key < 0: //插入右分支
		if root.Right == nil {
			var node = &TreeNode{
				Val:   key,
				Left:  nil,
				Right: nil,
			}
			root.Right = node
			return
		}
		insert(root.Right, key)
	case root.Val-key == 0: //不插入
		return
	}
}

func countNodes(root *TreeNode) int {
	//var cnt int
	//if root == nil {
	//	return cnt
	//}
	//cnt += 1
	//cnt += countNodes(root.Left)
	//cnt += countNodes(root.Right)
	//return cnt
	if root == nil {
		return 0
	}
	var stack = list.New()
	var cnt int
	stack.PushBack(root)
	for stack.Len() > 0 {
		ele := stack.Front()
		cnt++
		root = ele.Value.(*TreeNode)
		stack.Remove(ele)
		if root.Left != nil {
			stack.PushBack(root.Left)
		}
		if root.Right != nil {
			stack.PushBack(root.Right)
		}
	}
	return cnt
}
