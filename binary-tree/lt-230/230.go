package lt230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	rank, res := 0, 0

	traverse(root, k, &rank, &res)

	return res
}

func traverse(root *TreeNode, k int, rank, res *int) {
	if root == nil {
		return
	}

	traverse(root.Left, k, rank, res)

	*rank++
	if k == *rank {
		*res = root.Val
		return
	}

	traverse(root.Right, k, rank, res)
}
