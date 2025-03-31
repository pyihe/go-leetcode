package main

import (
	"fmt"
)

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
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
	fmt.Printf("%d\n", maxDepth(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxHeight(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
}

func maxHeight(h1, h2 int) int {
	if h1 > h2 {
		return h1
	}
	return h2
}

//levelNode 用于记录每一层的节点
func maxDepth(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}
	var levelNode = []*TreeNode{root}

	for len(levelNode) > 0 {
		sz := len(levelNode)
		for sz > 0 {
			node := levelNode[0]
			levelNode = levelNode[1:]
			if node.Left != nil {
				levelNode = append(levelNode, node.Left)
			}
			if node.Right != nil {
				levelNode = append(levelNode, node.Right)
			}
			sz--
		}
		result++
	}
	return result
}
