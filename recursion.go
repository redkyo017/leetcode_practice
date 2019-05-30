package main

import "log"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		a := s[i]
		s[i] = s[j]
		s[j] = a
		i++
		j--
	}
	log.Println(string(s))
}

var nodes4 = ListNode{4, nil}
var nodes3 = ListNode{3, &nodes4}
var nodes2 = ListNode{2, &nodes3}
var nodes1 = ListNode{1, &nodes2}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		log.Println("here")
		return head
	}
	// head, head.Next = head.Next, head
	n := head.Next
	head.Next = swapPairs(head.Next.Next)
	n.Next = head
	return n
}

func generate(numRows int) [][]int {
	result := [][]int{}
	for i := 1; i <= numRows; i++ {
		if i == 1 {
			result = append(result, []int{1})
			continue
		}
		r := make([]int, i)
		result = append(result, r)
		for j := 1; j <= i; j++ {
			if j == 1 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j+1]
			}
		}
	}
	log.Println(result)
	return result
}

func init() {
	generate(5)
}
