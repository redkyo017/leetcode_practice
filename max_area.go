package main

func isMax(height []int) int {
	// var rightside, leftside, rightindex, leftindex, maxArea int
	var maxArea int

	for i := 0; i < len(height)-1; i++ {
		rightside, rightindex := height[i], i
		for j := len(height) - 1; j > 0; j-- {
			leftside, leftindex := height[j], j

			bottomside := leftindex - rightindex
			var topside int
			if rightside > leftside {
				topside = leftside
			} else {
				topside = rightside
			}

			area := bottomside * topside
			// log.Println(rightside, leftside, bottomside, area)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
