package main

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回false
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//根结点的最小和最大的高度差不能超过1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	leftHeight := maxHeight(root.Left)
	rightHeight := maxHeight(root.Right)
	if leftHeight-rightHeight > 1 || leftHeight-rightHeight < -1 {
		return false
	}

	return true
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxHeight(root.Left), maxHeight(root.Right)) + 1
}

func max(h1, h2 int) int {
	if h1 > h2 {
		return h1
	}
	return h2
}
