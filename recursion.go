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
	for i := 0; i < numRows; i++ {
		if i == 0 {
			result = append(result, []int{1})
			continue
		}
		r := make([]int, i+1)
		result = append(result, r)
		// log.Println(result)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}
	log.Println(result)
	return result
}

func getRow(rowIndex int) []int {
	result := [][]int{}
	for i := 0; i <= rowIndex; i++ {
		if i == 0 {
			result = append(result, []int{1})
			continue
		}
		r := make([]int, i+1)
		result = append(result, r)
		// log.Println(result)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}
	log.Println(result[rowIndex])
	return result[rowIndex]
}

var lnode5 = ListNode{5, nil}
var lnode4 = ListNode{4, &lnode5}
var lnode3 = ListNode{3, &lnode4}
var lnode2 = ListNode{2, &lnode3}
var lnode1 = ListNode{1, &lnode2}

func reverseList(head *ListNode) *ListNode {
	// // ITERATIVE
	// prev := &ListNode{}
	// curr := head

	// for curr != nil {
	// 	temp := curr.Next
	// 	curr.Next = prev
	// 	prev = curr
	// 	curr = temp
	// }

	// return prev

	// RECURSIVE
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func fib(N int) int {
	if N < 2 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

func climb_stairs(i int, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return n
	}
	return climb_stairs(i+1, n) + climb_stairs(i+2, n)
}

var memo = map[int]int{}

func memoization_climb_stairs(i int, n int, m map[int]int) int {
	if i > n {
		return 0
	}
	if i == n {
		return n
	}
	v, ok := memo[i]
	if ok {
		return v
	}
	memo[i] = memoization_climb_stairs(i+1, n, memo) + memoization_climb_stairs(i+2, n, memo)
	return memo[i]
}

func climbStairs(n int) int {
	// BRUTE FROCE
	// res := climb_stairs(0, n)

	// MEMOIZATION
	// res := memoization_climb_stairs(0, n, memo)

	// DYNAMIC PROGRAMMING
	if n == 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}
	return second
	return 0
}

func init() {
	log.Println(climbStairs(6))
}
