package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	pad := 0
	origin := x
	for x != 0 {
		pop := x % 10
		x = x / 10
		pad = pad*10 + pop
	}
	if pad != origin {
		return false
	}
	return true
}
