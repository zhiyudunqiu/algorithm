package lt1373

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {
	maxSum := 0
	traverse(root, &maxSum)

	return maxSum
}

func traverse(root *TreeNode, maxSum *int) (bool, int, int, int) {
	if root == nil {
		return true, 4*10e4+1, -4*10e4-1, 0
	}

	leftIsBST, leftMin, leftMax, leftSum := traverse(root.Left, maxSum)
	rightIsBST, rightMin, rightMax, rightSum := traverse(root.Right, maxSum)

	if leftIsBST && rightIsBST && root.Val > leftMax && root.Val < rightMin {
		if *maxSum < leftSum + rightSum + root.Val {
			*maxSum = leftSum + rightSum + root.Val
		}

		return true, min(leftMin, root.Val), max(rightMax, root.Val), leftSum + rightSum + root.Val
	}

	return false, 4*10e4+1, -4*10e4-1, 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
