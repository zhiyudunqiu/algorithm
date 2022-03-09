package lt222

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 判断是否是完美二叉树
	lc, rc := 0, 0
	for leftSubTree := root; leftSubTree.Left != nil; leftSubTree = leftSubTree.Left {
		lc++
	}
	for rightSubTree := root; rightSubTree.Right != nil; rightSubTree = rightSubTree.Right {
		rc++
	}
	if lc == rc {
		return int(powN(float64(2), lc+1)) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func powN(x float64, n int) float64 {
	var res float64 = 1.0
	for n > 0 {
		res *= x
		n--
	}

	return res
}
