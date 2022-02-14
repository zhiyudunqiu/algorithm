package lt652

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	subTreesMap := make(map[string]int, 0)
	res := make([]*TreeNode, 0)

	traverse(root, subTreesMap, &res)

	return res
}

func traverse(root *TreeNode, subTreesMap map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}

	left := traverse(root.Left, subTreesMap, res)
	right := traverse(root.Right, subTreesMap, res)
	subTree := fmt.Sprintf("%s,%s,%d", left, right, root.Val)

	count, ok := subTreesMap[subTree]
	if ok && count == 1 {
		*res = append(*res, root)
	}
	subTreesMap[subTree] = count + 1

	return subTree
}
