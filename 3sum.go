package main

import "log"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	log.Println(nums)
	setNum := make(map[int]int)
	for _, num := range nums {
		setNum[num] = num
	}
	for i := 0; i < len(setNum); i++ {
		for j := i + 1; j < len(setNum); j++ {
			for k := j + 1; k < len(setNum); k++ {
				if setNum[i]+setNum[j]+setNum[k] == 0 {
					result = append(result, []int{setNum[i], setNum[j], setNum[k]})
				}
			}
		}
	}
	return result
}

// [-1, 0, 1, 2, -1, -4]
