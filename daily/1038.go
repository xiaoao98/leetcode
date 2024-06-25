package daily

func bstToGst(root *TreeNode) *TreeNode {
	nowToAdd := 0
	var rightInorder func(*TreeNode)
	rightInorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		rightInorder(node.Right)
		tmp := node.Val
		node.Val += nowToAdd
		nowToAdd += tmp
		rightInorder(node.Left)
	}
	rightInorder(root)
	return root
}