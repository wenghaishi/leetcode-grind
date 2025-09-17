package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {

}

func treeTraversal(root *TreeNode) ([]int) {
	result := []int{}

	var helper func(root *TreeNode)

	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Left)
		helper(root.Right)

		result = append(result, root.Val)

	}

	helper(root)

	return result
}