package main

import (
	"log"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for j := 0; j < len(nums)-3; j++ {
		if j == 0 || (j > 0 && nums[j] != nums[j-1]) || nums[j] == 0 {
			for i := j + 1; i < len(nums)-2; i++ {
				// if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
				log.Println(j, i, nums[j], nums[i])
				if nums[i] == 0 || nums[i] != nums[i-1] {
					lo := i + 1
					hi := len(nums) - 1
					sum := target - (nums[i] + nums[j])
					log.Println(nums[i], nums[j], sum)

					for lo < hi {
						if nums[lo]+nums[hi] == sum {
							log.Println(nums[i], nums[j], nums[lo], nums[hi])
							result = append(result, []int{nums[j], nums[i], nums[lo], nums[hi]})
							// for lo < hi && nums[lo] == nums[lo+1] {
							// 	lo++
							// }
							// for lo < hi && nums[hi] == nums[hi-1] {
							// 	hi--
							// }
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
