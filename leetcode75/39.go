package leetcode75

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ret *TreeNode
	_ = dfs(root, p, q, &ret)
	return ret
}

func dfs(node, p, q *TreeNode, ret **TreeNode) int {
	if node == nil {
		return 0
	}
	if *ret != nil {
		return 0
	}
	leftValue := dfs(node.Left, p, q, ret)
	rightValue := dfs(node.Right, p, q, ret)
	nodeValue := 0
	if node == p || node == q {
		nodeValue = 1
	}
	if leftValue+rightValue+nodeValue >= 2 {
		*ret = node
		return 0
	}
	return leftValue + rightValue + nodeValue
}