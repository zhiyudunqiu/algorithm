package lt116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 迭代解法-2
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// 按照完美二叉树的左边遍历
	for leftFirthNode := root; leftFirthNode.Left != nil; leftFirthNode = leftFirthNode.Left {
		// 遍历当前层次，并且调整下一层节点的指针；由于root第一层相当于初始化好的，可以正常遍历
		for node := leftFirthNode; node != nil; node = node.Next {
			node.Left.Next = node.Right // 调整node子节点的next

			// 调整node和node.Next中相邻子节点的next
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	return root
}

// 迭代解法-1
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	queue := []*Node{root}
// 	nextQueue := []*Node{}
// 	for len(queue) > 0 {
// 		for i, n := range queue {
// 			if i+1 < len(queue) {
// 				queue[i].Next = queue[i+1]
// 			}

// 			if n.Left != nil {
// 				nextQueue = append(nextQueue, n.Left)
// 			}

// 			if n.Right != nil {
// 				nextQueue = append(nextQueue, n.Right)
// 			}
// 		}

// 		queue = nextQueue
// 		nextQueue = nil
// 	}

// 	return root
// }

// 递归解法
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	connectTwoNodes(root.Left, root.Right)
// 	return root
// }

// func connectTwoNodes(leftNode, rightNode *Node) {
// 	if leftNode == nil || rightNode == nil {
// 		return
// 	}

// 	leftNode.Next = rightNode

// 	connectTwoNodes(leftNode.Left, leftNode.Right)
// 	connectTwoNodes(leftNode.Right, rightNode.Left)
// 	connectTwoNodes(rightNode.Left, rightNode.Right)
// }
