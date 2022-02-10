package lt105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}

	leftTreeLen := index
	root.Left = buildTree(preorder[1:leftTreeLen+1], inorder[:index])

	if leftTreeLen < len(preorder) - 1 {
		root.Right = buildTree(preorder[leftTreeLen+1:], inorder[index+1:])
	}
	
	return root
}
