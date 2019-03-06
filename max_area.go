package main

import "log"

func isMax(height []int) int {
	var rightside, leftside, rightindex, leftindex, maxArea int

	for i := 0; i < len(height)/2; i++ {
		for j := len(height) - 1; j > len(height)/2; j-- {
			if height[i] > rightside {
				rightside = height[i]
				rightindex = i
			}

			if height[j] > leftside {
				leftside = height[j]
				leftindex = j
			}

			bottomside := leftindex - rightindex
			var topside int

			if rightside > leftside {
				topside = leftside
			} else {
				topside = rightside
			}

			area := bottomside * topside

			log.Println(rightside, leftside, bottomside, area)

			if area > maxArea {
				maxArea = area
			}
		}

	}

	return maxArea
}
