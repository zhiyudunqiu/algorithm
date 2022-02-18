package lt297

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#,"
	}

	res := fmt.Sprintf("%d,", root.Val)
	res += this.serialize(root.Left)
	res += this.serialize(root.Right)

	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(strings.Trim(data, ","), ",")

	var build func() *TreeNode
	build = func() *TreeNode {
		if len(nodes) == 0 {
			return nil
		}

		if nodes[0] == "#" {
			nodes = nodes[1:]
			return nil
		}

		nodeVal, _ := strconv.Atoi(nodes[0])
		nodes = nodes[1:]

		return &TreeNode{nodeVal, build(), build()}
	}

	return build()
}

// func deserializeTool(nodes *[]string) *TreeNode {
// 	// 空列表返回nil
// 	if len(*nodes) == 0 {
// 		return nil
// 	}

// 	// 节点为空节点
// 	if (*nodes)[0] == "#" {
// 		// 弹出根节点
// 		*nodes = (*nodes)[1:]
// 		return nil
// 	}

// 	// 生成根节点
// 	nodeVal, err := strconv.Atoi((*nodes)[0])
// 	if err != nil {
// 		panic(fmt.Sprintf("convert string to int error: %v", err.Error()))
// 	}
// 	root := &TreeNode{Val: nodeVal}

// 	// 弹出根节点
// 	*nodes = (*nodes)[1:]

// 	// 递归处理左右子树
// 	root.Left = deserializeTool(nodes)
// 	root.Right = deserializeTool(nodes)

// 	return root
// }
