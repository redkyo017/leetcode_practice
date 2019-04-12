package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

var node5 = ListNode{5, nil}
var node4 = ListNode{4, &node5}
var node3 = ListNode{3, &node4}
var node2 = ListNode{2, &node3}
var node1 = ListNode{1, &node2}

var node10 = ListNode{10, nil}
var node9 = ListNode{8, &node10}
var node8 = ListNode{6, &node9}
var node7 = ListNode{4, &node8}
var node6 = ListNode{2, &node7}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	dump := result
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			result.Next = l2
			l2 = l2.Next
		} else {
			result.Next = l1
			l1 = l1.Next
		}
		// if result.Next != nil {
		result = result.Next
		// }
	}
	if l1 != nil {
		result.Next = l1
	}
	if l2 != nil {
		result.Next = l2
	}
	// log.Println("con co", l1, l2, result)

	// node := dump.Next
	// for node != nil {
	// 	log.Println("con heo", node)
	// 	node = node.Next
	// }

	return dump.Next
}
