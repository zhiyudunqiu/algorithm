// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
package lt104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil { // 递归退出条件
		return 0
	}

	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)

	if leftMaxDepth > rightMaxDepth {
		return 1 + leftMaxDepth
	}

	return 1 + rightMaxDepth
}
