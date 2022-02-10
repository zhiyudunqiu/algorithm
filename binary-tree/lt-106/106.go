package lt106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			index = i
			break
		}
	}

	leftTreeLen := index
	root.Left = buildTree(inorder[:index], postorder[:index])
	if leftTreeLen < len(inorder) - 1 {
		root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	}
	
	return root
}
