package daily

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode, 0)
	childSet := make(map[int]bool, 0)
	for _, desc := range descriptions {
		parent := desc[0]
		child := desc[1]
		isLeft := desc[2]
		var parentNode *TreeNode
		if _, exists := nodeMap[parent]; !exists {
			parentNode = &TreeNode{Val: parent}
			nodeMap[parent] = parentNode
		} else {
			parentNode = nodeMap[parent]
		}
		var childNode *TreeNode
		if _, exists := nodeMap[child]; !exists {
			childNode = &TreeNode{Val: child}
			nodeMap[child] = childNode
		} else {
			childNode = nodeMap[child]
		}
		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
		childSet[child] = false
	}
	for key, val := range nodeMap {
		if _, exists := childSet[key]; !exists {
			return val
		}
	}
	return nil
}