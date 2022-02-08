package lt144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	preorder(root, &res)

	return res
}

func preorder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}

	*list = append(*list, root.Val) // 对于nil slice，append内部会初始化nil slice
	preorder(root.Left, list)
	preorder(root.Right, list)
}
