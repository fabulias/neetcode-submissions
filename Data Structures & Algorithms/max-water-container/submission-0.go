func maxArea(heights []int) int {
	n := len(heights)
	left, right := 0, n-1
	mostWater := 0
	for left < right {
		water := (right - left) * min(heights[right], heights[left]) 
		
		mostWater = max(water, mostWater)
		fmt.Println(mostWater, water)
		if heights[right] < heights[left] {
			right--
		} else {
			left++
		}
	} 
	return mostWater
}
