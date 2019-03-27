package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// var result int

	sort.Ints(nums)

	// var set [3]int
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		lo := i + 1
		hi := len(nums) - 1

		for lo < hi {
			sum := (nums[i] + nums[lo] + nums[hi])

			if math.Abs(float64(target-closest)) > math.Abs(float64(target-sum)) {
				closest = sum
				if closest == target {
					return closest
				}
			}
			if sum > target {
				hi--
			} else {
				lo++
			}
		}
	}

	return closest
}

// [-1, 2, 1, -4], 1 : 2
// [0,0,0], 1: 0
// [1,1,1,0], -100: 2
// [1,1,1,0], 100: 3
// [0,2,1,-3], 1: 0
// [1,1,-1,-1,3], -1: -1
