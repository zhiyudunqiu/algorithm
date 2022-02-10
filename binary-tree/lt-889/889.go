package lt889

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	root := &TreeNode{Val: preorder[0]}
	
	index := 0
	for i := 0; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			index = i
			break
		}
	}

	// 当作左子树处理
	leftTreeLen := index + 1

	// 至少有两个元素的时候，左子树根节点不可能是最后一个元素，所以leftTreeLen+1不会数组越界
	root.Left = constructFromPrePost(preorder[1:leftTreeLen+1], postorder[:index+1])
	root.Right = constructFromPrePost(preorder[leftTreeLen+1:], postorder[index+1:len(postorder)-1])

	return root
}
