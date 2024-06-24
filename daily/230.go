package daily

func kthSmallest(root *TreeNode, k int) int {
	var ret int
	var inorder1 func(*TreeNode)
	inorder1 = func(node *TreeNode) {
		if node == nil || k <= 0 {
			return
		}
		inorder1(node.Left)
		k--
		if k == 0 {
			ret = node.Val
		}
		inorder1(node.Right)
	}
	inorder1(root)
	return ret
}