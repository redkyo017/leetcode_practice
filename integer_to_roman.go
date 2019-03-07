package main

import (
	"log"
	"strconv"
)

func intToRoman(num int) string {
	m := map[int]string{
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
	log.Println(m)

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
		log.Println(trunkNum/1000, trunkNum%1000)
	}
	return ""
}
