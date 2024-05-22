package daily

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	ret := 0
	var helper func(*TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftVal := helper(node.Left)
		rightVal := helper(node.Right)
		ret += int(math.Abs(float64(leftVal))) + int(math.Abs(float64(rightVal)))
		return leftVal + rightVal + node.Val - 1
	}
	helper(root)
	return ret
}