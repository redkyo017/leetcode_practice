package main

import "log"

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
var node2 = ListNode{2, nil}
var node1 = ListNode{1, &node2}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dum := ListNode{}
	dum.Next = head
	length := 0
	// next := head.Next
	first := head
	for first != nil {
		length++
		first = first.Next
	}

	if length == 1 && n == 1 {
		return nil
	}

	length -= n

	if length == 0 {
		return head.Next
	}

	first = &dum
	log.Println("hehe", length)
	for length > 0 {
		length--
		first = first.Next
	}
	first.Next = first.Next.Next

	node := dum.Next
	for node != nil {
		log.Println("con heo", node)
		node = node.Next
	}
	// return first.Next
	return head
}
