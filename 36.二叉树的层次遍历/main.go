package main

import "fmt"

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

var (
	root = &TreeNode{
		Val: 1,
	}
)

func init() {
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	leftLeft := &TreeNode{Val: 4}
	rightRight := &TreeNode{Val: 5}

	left.Left = leftLeft
	right.Right = rightRight

	root.Left = left
	root.Right = right
}

func main() {
	fmt.Printf("%v\n", levelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var levelNode = []*TreeNode{root}
	var levelResult []int
	for len(levelNode) > 0 {
		sz := len(levelNode)
		for sz > 0 {
			node := levelNode[0]
			levelNode = levelNode[1:]
			levelResult = append(levelResult, node.Val)
			if node.Left != nil {
				levelNode = append(levelNode, node.Left)
			}
			if node.Right != nil {
				levelNode = append(levelNode, node.Right)
			}
			sz--
		}
		result = append(result, levelResult)
		levelResult = []int{}
	}
	return result
}
