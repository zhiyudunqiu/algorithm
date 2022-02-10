package lt226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	tmpNode := root.Left
	root.Left = root.Right
	root.Right = tmpNode

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}