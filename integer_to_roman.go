package main

import (
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	m := map[int]string{
		0:    "",
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	result := []string{}
	numString := strconv.Itoa(num)
	for len(numString) > 0 {
		trunk := ""
		if len(numString) > 4 {
			trunk = numString[:4]
			numString = numString[4:]
		} else {
			trunk = numString[:]
			numString = ""
		}

		trunkNum, _ := strconv.Atoi(trunk)
		str := ""

		mod := trunkNum % 10
		str = m[mod] + str

		trunkNum = trunkNum / 10
		if trunkNum <= 0 {
			result = append(result, str)
			continue
		}

		mod = trunkNum % 10
		if mod == 9 {
			str = "XC" + str
		} else if mod == 5 {
			str = "L" + str
		} else {
			str = m[mod] + "X" + str
		}

		trunkNum = trunkNum / 10
		if trunkNum <= 0 {
			result = append(result, str)
			continue
		}

		mod = trunkNum % 10
		if mod == 9 {
			str = "CM" + str
		} else if mod == 5 {
			str = "D" + str
		} else {
			str = m[mod] + "C" + str
		}

		trunkNum = trunkNum / 10
		if trunkNum <= 0 {
			result = append(result, str)
			continue
		} else if trunkNum == 1 {
			str = "M" + str
		} else {
			str = m[trunkNum] + "M" + str
		}

		result = append(result, str)
	}
	return strings.Join(result, ",")
}
