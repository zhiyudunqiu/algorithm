package lt543

// https://leetcode-cn.com/problems/diameter-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	maxDepth(root, &diameter)

	return diameter
}

func maxDepth(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	leftMaxDepth := maxDepth(root.Left, diameter)
	rightMaxDepth := maxDepth(root.Right, diameter)

	if leftMaxDepth + rightMaxDepth > *diameter {
		*diameter = leftMaxDepth + rightMaxDepth
	}

	if leftMaxDepth > rightMaxDepth {
		return 1 + leftMaxDepth
	}

	return 1 + rightMaxDepth
}
