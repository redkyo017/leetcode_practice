package main

import "log"

func backtrack(s string, left int, right int, ans []string, n int) []string {
	if len(s) == 2*n {
		ans = append(ans, s)
		return ans
	}
	if left < n {
		openp := "("
		return backtrack(s+openp, left+1, right, ans, n)
	}
	if right < left {
		closep := ")"
		return backtrack(s+closep, left, right+1, ans, n)
	}
	return ans
}

func generateParenthesis(n int) []string {
	result := []string{}

	result = backtrack("", 0, 0, result, n)
	log.Println(result)

	return result
}

// 2
// [
//   "(())",
//   "()()",
// ]

// 3
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

// 4
// [
//   "(((())))",
//   "(()()())",
//   "((()()))",
//   "((()))()",
//   "()((()))",
//   "()(())()",
//   "(())()()",
//   "()()(())",
//   "(())(())",
//   "()()()()"
// ]
