package lt95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return build(1, n)
}

func build(low, hight int) []*TreeNode {
	if low > hight {
		return []*TreeNode{nil}
	}

	res := []*TreeNode{}
	for i := low; i <= hight; i++ {
		lefts := build(low, i-1)
		rights := build(i+1, hight)

		for _, l := range lefts {
			for _, r := range rights {
				root := &TreeNode{
					Val: i,
					Left: l,
					Right: r,
				}

				res = append(res, root)
			}
		}
	}

	return res
}
