package main

import (
	"strconv"
)

func myAtoi(str string) int {
	if len(str) <= 0 {
		return 0
	}

	if str[0] != 32 && (str[0] < 48 || str[0] > 57) && (string(str[0]) != "-" && string(str[0]) != "+") {
		return 0
	}
	var neg, pos bool
	for k, val := range str {
		if val != 32 {
			str = str[k:]
			break
		}
	}

	if (str[0] < 48 || str[0] > 57) && (string(str[0]) != "-" && string(str[0]) != "+") {
		return 0
	}

	if string(str[0]) == "-" {
		if len(str) == 1 {
			return 0
		}
		neg = true
		str = str[1:]
	}

	if string(str[0]) == "+" {
		if len(str) == 1 {
			return 0
		}
		pos = true
		str = str[1:]
	}
	if neg && pos {
		return 0
	}

	for k, val := range str {
		// if (val >= 48 && val <= 57) || (string(val) == "-" || string(val) == "+") {
		if val >= 48 && val <= 57 {
			continue
		} else {
			str = str[:k]
			break
		}
	}

	n := 0
	MAX := 1<<31 - 1
	MIN := -1 << 31

	for _, val := range []byte(str) {
		if n > MAX || n < MIN {
			break
		}
		if val < 48 || val > 57 {
			return 0
		}

		num, _ := strconv.ParseInt(string(val), 10, 64)
		n = n*10 + int(num)
	}

	if neg {
		n = -n
	}

	if n > MAX {
		return MAX
	}
	if n < MIN {
		return MIN
	}

	return n
}

// func main() {
// 	test := "   - 321"
// 	// test := "  0000000000012345678"
// 	// test := "+"
// 	// test := "   -42"
// 	// test := "-5-"
// 	// test := "-91283472332"

// 	// tetNum, _ := strconv.Atoi(test)
// 	result := myAtoi(test)
// 	log.Println(result)

// }
