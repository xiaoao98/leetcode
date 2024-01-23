package leetcode75

var longest int = 0

func longestZigZag(root *TreeNode) int {
	longest = 0
	_ = findLonestZigZag(root.Left, true)
	_ = findLonestZigZag(root.Right, false)
	return longest
}

func findLonestZigZag(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}
	if isLeft {
		_ = findLonestZigZag(node.Left, true)
		rightLongest := 1 + findLonestZigZag(node.Right, false)
		if longest < rightLongest {
			longest = rightLongest
		}
		return rightLongest
	} else {
		_ = findLonestZigZag(node.Right, false)
		leftLongest := 1 + findLonestZigZag(node.Left, true)
		if longest < leftLongest {
			longest = leftLongest
		}
		return leftLongest
	}
}
