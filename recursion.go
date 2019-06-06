package main

import (
	"log"
	"math"
)

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

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n == 1 {
		return x
	}
	if x == 0 {
		return 0
	}

	neg := false
	// o := x
	if n < 0 {
		n = int(math.Abs(float64(n)))
		neg = true
	}
	// for i := 2; i <= n; i++ {
	// 	x *= o
	// }
	var result float64 = 1
	for n > 0 {
		if n%2 == 0 {
			x *= x
			n /= 2
		} else {
			n = n - 1
			result = result * x
			n /= 2
			x *= x
		}
	}
	if neg {
		result = 1 / result
	}
	return result
}

var lnodeR6 = ListNode{4, nil}
var lnodeR5 = ListNode{3, &lnodeR6}
var lnodeR4 = ListNode{1, &lnodeR5}

var lnodeR3 = ListNode{4, nil}
var lnodeR2 = ListNode{2, &lnodeR3}
var lnodeR1 = ListNode{1, &lnodeR2}

func mergeTwoListsRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		result = l2
		result.Next = mergeTwoListsRecursion(l1, l2.Next)
	} else {
		result = l1
		result.Next = mergeTwoListsRecursion(l1.Next, l2)
	}

	return result
}

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	if K%2 == 0 {
		n := kthGrammar(N-1, K/2)
		if n == 0 {
			return 1
		} else {
			return 0
		}
	} else {
		return kthGrammar(N-1, (K+1)/2)
	}

	return 0
}

func init() {
	// 1.00000 -2147483648
	// log.Println(myPow(2.00000, 10))

	// node := mergeTwoListsRecursion(&lnodeR1, &lnodeR4)
	// for node != nil {
	// 	log.Println(node.Val)
	// 	node = node.Next
	// }

	log.Println(kthGrammar(4, 8))
}
