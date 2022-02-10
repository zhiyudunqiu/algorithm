package lt654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
		return nil
	}

	index := 0
	maxNum := nums[index]

	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNum {
			index = i
			maxNum = nums[i]
		}
	}

	root := &TreeNode{Val: maxNum}
	root.Left = constructMaximumBinaryTree(nums[:index])
	
	if index < len(nums)-1 {
		root.Right = constructMaximumBinaryTree(nums[index+1:])
	}
	
	return root
}