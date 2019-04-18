package main

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

func dfs(left, right, idx int, bytes []byte, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}

	if left > 0 {
		bytes[idx] = '('
		dfs(left-1, right, idx+1, bytes, res)
	}

	if right > 0 && left < right {
		bytes[idx] = ')'
		dfs(left, right-1, idx+1, bytes, res)
	}
}

func generateParenthesis(n int) []string {
	// result := []string{}

	// result = backtrack("", 0, 0, result, n)
	// log.Println(result)

	// return result
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)
	dfs(n, n, 0, bytes, &res)
	return res
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
