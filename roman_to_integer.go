package main

import (
	"strings"
)

func romanToInt(s string) int {
	M := map[string]int{"M": 1000, "MM": 2000, "MMM": 3000, "C": 100, "CC": 200, "CCC": 300, "CD": 400, "D": 500, "DC": 600, "DCC": 700, "DCCC": 800, "CM": 900, "X": 10, "XX": 20, "XXX": 30, "XL": 40, "L": 50, "LX": 60, "LXX": 70, "LXXX": 80, "XC": 90, "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9}

	a := strings.Split(s, "")
	res := 0
	var m int

	for i, _ := range a {
		str := strings.Join(a[m:i+1], "")
		if v, ok := M[str]; ok {
			if i+1 == len(a) {
				res += v
				break
			} else {
				continue
			}
		}

		roman := a[m:i]
		key := strings.Join(roman, "")
		res += M[key]
		m = i

		if i == len(a)-1 {
			if v, ok := M[a[i]]; ok {
				res += v
				break
			}
		}
	}

	return res
}

// LVIII
// MCMXCIV
// DCXXI
// III
// IV
// IX
