package daily

func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		helper(node.Right)
		ret = append(ret, node.Val)
	}
	helper(root)
	return ret
}