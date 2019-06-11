package main

import "log"

func helperValid(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	val := root.Val
	log.Println(lower, val, upper)
	if val <= lower || val >= upper {
		return false
	}
	if !helperValid(root.Left, lower, val) {
		return false
	}
	if !helperValid(root.Right, val, upper) {
		return false
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	// MAX := 1 << 31
	MIN := -1<<31 - 1
	// APPROACH 1: RECURSION
	// return helperValid(root, MIN, MAX)

	// APPROACH 2: ITERATION
	// if root == nil {
	// 	return true
	// }
	// type stackNode struct {
	// 	node  *TreeNode
	// 	lower int
	// 	upper int
	// }

	// var stack []stackNode
	// stack = append(stack, stackNode{root, MIN, MAX})
	// for len(stack) > 0 {
	// 	pop := stack[0]
	// 	stack = stack[1:]
	// 	if pop.node == nil {
	// 		continue
	// 	}
	// 	val := pop.node.Val
	// 	if val <= pop.lower || val >= pop.upper {
	// 		return false
	// 	}
	// 	stack = append(stack, stackNode{pop.node.Left, pop.lower, val})
	// 	stack = append(stack, stackNode{pop.node.Right, val, pop.upper})
	// }

	// APPROACH 3 : Inorder traversal
	var stack []*TreeNode
	var inorder int = MIN
	for len(stack) >= 0 && root != nil {
		log.Println("con co")
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root := stack[0]
		stack = stack[1:]
		if root.Val <= inorder {
			log.Println(root, root.Val, inorder)
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

// [2,1,3]
// [5,1,4,null,null,3,6]
// [10,5,15,null,null,6,20]
// [2147483647]

var tNode1 = TreeNode{5, nil, nil}
var tNode2 = TreeNode{2, &tNode1, &tNode3}
var tNode3 = TreeNode{6, nil, nil}
var tNode4 = TreeNode{15, &tNode3, &tNode6}
var tNode5 = TreeNode{10, &tNode1, &tNode4}
var tNode6 = TreeNode{20, nil, nil}

func init() {
	log.Println(isValidBST(&tNode5))
}
