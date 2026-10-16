# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def createBinaryTree(self, descriptions: list[list[int]]) -> TreeNode | None:
        dic = {}
        childSet = set()
        for descrip in descriptions:
            par = descrip[0]
            child = descrip[1]
            childSet.add(child)
            parNode = None
            if par in dic:
                parNode = dic[par]
            else:
                parNode = TreeNode(par)
                dic[par] = parNode
            childNode = None
            if child in dic:
                childNode = dic[child]
            else:
                childNode = TreeNode(child)
                dic[child] = childNode
            if descrip[2] == 1:
                parNode.left = childNode
            else:
                parNode.right = childNode
        for descrip in descriptions:
            if descrip[0] not in childSet:
                return dic[descrip[0]]
        return None