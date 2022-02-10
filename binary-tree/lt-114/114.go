package lt114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	if root == nil { // 边界条件
		return
	}

	flatten(root.Left) // 拉直左子树
	flatten(root.Right) // 拉直右子树

	// 将左子树列表移到右子树节点上，并将左子树节点只为nil
	rightList := root.Right
	root.Right = root.Left
	root.Left = nil

	// 将原右子树，移到现右子树的末尾，注意右子树为空的情况
	node := root
	for node.Right != nil {
		node = node.Right
	}
	node.Right = rightList
}