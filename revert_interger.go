package main

import "math"

// MY SOLUTIONS
// func reverse(n int) int {
// 	var isNegative bool
// 	if n < 0 {
// 		n = int(math.Abs(float64(n)))
// 		isNegative = true
// 	}
// 	nStr := []byte(strconv.Itoa(n))

// 	for i := 0; i < len(nStr)/2; i++ {
// 		j := len(nStr) - 1 - i
// 		nStr[i], nStr[j] = nStr[j], nStr[i]
// 	}
// 	result := string(nStr)
// 	if isNegative {
// 		result = "-" + result
// 	}
// 	r, _ := strconv.Atoi(result)
// 	return r
// }

// RIGHT SOLUTION

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x = x / 10

		if rev > math.MaxInt64/10 || (rev == math.MaxInt64/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt64/10 || (rev == math.MinInt64/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}

	return rev
}
