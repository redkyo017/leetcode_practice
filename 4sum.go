package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			for j := len(nums) - 1; j > 2; j-- {
				if j == len(nums)-1 || (j < len(nums)-1 && nums[j] != nums[j+1]) {
					lo := i + 1
					hi := j - 1
					sum := target - (nums[i] + nums[j])
					for lo < hi {
						if nums[lo]+nums[hi] == sum {
							// log.Println(nums)
							// log.Println(i, lo, hi, j)
							// log.Println(nums[i], nums[lo], nums[hi], nums[j])
							// log.Println("---------")
							result = append(result, []int{nums[i], nums[lo], nums[hi], nums[j]})
							for lo < hi && nums[lo] == nums[lo+1] {
								lo++
							}
							for lo < hi && nums[hi] == nums[hi-1] {
								hi--
							}
							lo++
							hi--
						} else if nums[lo]+nums[hi] < sum {
							lo++
						} else {
							hi--
						}
					}
				}
			}
		}
	}
	// log.Println("length:", len(result))
	return result
}

// [1, 0, -1, 0, -2, 2], 0
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]

// [0,0,0,0], 0
// [[0,0,0,0]]

// [-3,-2,-1,0,0,1,2,3], 0
// [
// 	[-3,-2,2,3],
// 	[-3,-1,1,3],
// 	[-3,0,0,3],
// 	[-3,0,1,2],
// 	[-2,-1,0,3],
// 	[-2,-1,1,2],
// 	[-2,0,0,2],
// 	[-1,0,0,1]
// ]

// [-5,5,4,-3,0,0,4,-2], 4
// [
//	[-5,0,4,5],
//	[-3,-2,4,5]
// ]

// [5,5,3,5,1,-5,1,-2], 4
// [[-5,1,3,5]]

// [-1,-5,-5,-3,2,5,0,4], -7
// [
//	[-5,-5,-1,4],
//	[-5,-3,-1,2]
// ]
