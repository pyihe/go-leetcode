package main

import (
	"container/list"
	"fmt"
)

/*
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的key对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7
*/

func main() {
	root := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	//[2,0,33, 1,25,40,  11,31,34,45,10,18,29,32, 36,43,46,4, 12,24,26,30,  35,39,42,44, 48,3,9, 14,22,  27,    38, 41,   47,49,  5, 13,15,21,23, 28,37,        8,   17,19,       7, 16,  20,6]
	//33
	num := []int{0, 33, 1, 25, 40, 11, 31, 34, 45, 10, 18, 29, 32, 36, 43, 46, 4, 12, 24, 26, 30, 35, 39, 42, 44, 48, 3, 9, 14, 22, 27, 38, 41, 47, 49, 5, 13, 15, 21, 23, 28, 37, 8, 17, 19, 7, 16, 20, 6}
	for _, v := range num {
		insert(root, v)
	}
	fmt.Printf("before: %v\n", root)
	deleteNode(root, 33)
	fmt.Printf("after: %v\n", root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tree *TreeNode) String() (result string) {
	treeList := list.New()
	treeList.PushBack(tree)

	for treeList.Len() > 0 {
		element := treeList.Front()
		node := element.Value.(*TreeNode)

		result += fmt.Sprintf("%d->", node.Val)
		treeList.Remove(element)

		if node.Left != nil {
			treeList.PushBack(node.Left)
		}
		if node.Right != nil {
			treeList.PushBack(node.Right)
		}
	}
	return
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

//my solution
func deleteNode(root *TreeNode, key int) *TreeNode {
	var pre *TreeNode
	var guard = root

	for guard != nil {
		if guard.Val > key {
			pre = guard
			guard = guard.Left
			continue
		}
		if guard.Val < key {
			pre = guard
			guard = guard.Right
			continue
		}

		//删除的不是叶子节点
		//1. 如果右孩子不为空，则将右子树中的最小值与当前节点替换，然后删除被替换的最小值叶子节点
		if guard.Right != nil {
			minPre, min := minNode(guard.Right)
			guard.Val = min.Val
			//删除max
			if minPre != nil {
				minPre.Left = min.Right
			} else {
				guard.Right = min.Right
			}
			return root
		}
		//2. 如果左孩子不为空，则将左子树中的最大值与当前节点替换，然后删除被替换的最大值叶子节点
		if guard.Left != nil {
			maxPre, max := maxNode(guard.Left)
			guard.Val = max.Val
			if maxPre != nil {
				maxPre.Right = max.Left
			} else {
				guard.Left = max.Left
			}
			return root
		}
		//删除的是叶子节点
		if pre == nil { //只有根结点，则直接删除，返回nil
			return nil
		} else { //删除对应叶子节点即可
			if pre.Left == guard {
				pre.Left = nil
			} else {
				pre.Right = nil
			}
			return root
		}
	}
	return root
}

func maxNode(root *TreeNode) (*TreeNode, *TreeNode) {
	var pre *TreeNode
	node := root
	for node.Right != nil {
		pre = node
		node = node.Right
	}
	return pre, node
}

func minNode(root *TreeNode) (*TreeNode, *TreeNode) {
	var pre *TreeNode
	node := root
	for node.Left != nil {
		pre = node
		node = node.Left
	}
	return pre, node
}

/**************official solution******************/
func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	//到这里意味已经查找到目标
	if root.Right == nil {
		//右子树为空
		return root.Left
	}
	if root.Left == nil {
		//左子树为空
		return root.Right
	}
	minNode := root.Right
	for minNode.Left != nil {
		//查找后继
		minNode = minNode.Left
	}
	root.Val = minNode.Val
	root.Right = deleteMinNode(root.Right)
	return root
}

func deleteMinNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteMinNode(root.Left)
	return root
}
