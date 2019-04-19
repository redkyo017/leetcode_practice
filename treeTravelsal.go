package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var treenode5 = TreeNode{7, nil, nil}
var treenode4 = TreeNode{15, nil, nil}
var treenode3 = TreeNode{20, &treenode4, &treenode5}
var treenode2 = TreeNode{9, nil, nil}
var rootNode = TreeNode{3, &treenode2, &treenode3}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		result = append(result, root.Val)

		if root.Left != nil {
			result = append(result, preorderTraversal(root.Left)...)
		}
		if root.Right != nil {
			result = append(result, preorderTraversal(root.Right)...)
		}
	}

	return result
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {

		if root.Left != nil {
			result = append(result, inorderTraversal(root.Left)...)
		}

		if root.Right != nil {
			result = append(result, inorderTraversal(root.Right)...)
		}

		result = append(result, root.Val)
	}

	return result
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {

		if root.Left != nil {
			result = append(result, postorderTraversal(root.Left)...)
		}

		if root.Right != nil {
			result = append(result, postorderTraversal(root.Right)...)
		}

		result = append(result, root.Val)
	}

	return result
}

// level - order Traversal

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	Q := []*TreeNode{}
	Q = append(Q, root)
	for len(Q) > 0 {
		lenQ := len(Q)
		currentlistVal := []int{}
		for i := 0; i < lenQ; i++ {
			currentNode := Q[0]
			if currentNode != nil {
				currentlistVal = append(currentlistVal, currentNode.Val)
				if currentNode.Left != nil {
					Q = append(Q, currentNode.Left)
				}
				if currentNode.Right != nil {
					Q = append(Q, currentNode.Right)
				}
			}
			Q = Q[1:]
		}

		result = append(result, currentlistVal)
	}

	return result
}
