package main

import "log"

func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	isSubtreeValid := isValidBST(root.Left) && isValidBST(root.Right)

	isValid := true
	if (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Right.Val <= root.Val) {
		isValid = false
	}
	if !isValid {
		return isValid
	}

	return isSubtreeValid && isValid
}

// [2,1,3]
// [5,1,4,null,null,3,6]
// [10,5,15,null,null,6,20]

var tNode1 = TreeNode{5, nil, nil}
var tNode2 = TreeNode{2, nil, nil}
var tNode3 = TreeNode{6, nil, nil}
var tNode4 = TreeNode{15, &tNode3, &tNode6}
var tNode5 = TreeNode{10, &tNode1, &tNode4}
var tNode6 = TreeNode{20, nil, nil}

func init() {
	log.Println(isValidBST(&tNode5))
}
