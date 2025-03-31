package main

/*
给定二叉树根结点root，此外树的每个结点的值要么是 0，要么是 1。

返回移除了所有不包含 1 的子树的原二叉树。

( 节点 X 的子树为 X 本身，以及所有 X 的后代。)

说明:
给定的二叉树最多有 100 个节点。
每个节点的值只会为 0 或 1 。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	return prune(root)
}

//如果节点左右孩子为空，且val==0，则剪掉，
func prune(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = prune(root.Left)
	root.Right = prune(root.Right)
	if root.Right == nil && root.Left == nil && root.Val == 0 {
		return nil
	}
	return root
}
