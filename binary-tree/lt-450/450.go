package lt450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key == root.Val {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		rightMinNode := root.Right
		for rightMinNode.Left != nil {
			rightMinNode = rightMinNode.Left
		}

		newNode := *rightMinNode // 注意：要进行一次值复制
		newNode.Left = root.Left
		newNode.Right = root.Right
		root = &newNode

		root.Right = deleteNode(root.Right, rightMinNode.Val)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
