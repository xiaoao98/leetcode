package leetcode75

func goodNodes(root *TreeNode) int {
	return findGoodNodes(root, -10000)
}

func findGoodNodes(node *TreeNode, comp int) int {
	if node == nil {
		return 0
	}
	if node.Val >= comp {
		return 1 + findGoodNodes(node.Left, node.Val) + findGoodNodes(node.Right, node.Val)
	} else {
		return findGoodNodes(node.Left, comp) + findGoodNodes(node.Right, comp)
	}
}