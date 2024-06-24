package daily

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    size := len(preorder)
	return buildTreeWrapper(preorder, inorder, 0, size-1, 0, size-1)
}

func buildTreeWrapper(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	if preStart == preEnd {
		return &TreeNode{Val: preorder[preStart]}
	}
	rootVal := preorder[preStart]
	rootPos := 0
	for i := inStart; i <= inEnd; i ++ {
		if inorder[i] == rootVal {
			rootPos = i
		}
	}
	leftNode := buildTreeWrapper(preorder, inorder, preStart+1, preStart+rootPos-inStart, inStart, rootPos-1)
	rightNode := buildTreeWrapper(preorder, inorder, preStart+rootPos-inStart+1, preEnd, rootPos+1, inEnd)
	return &TreeNode{Val: rootVal, Left: leftNode, Right: rightNode}
}