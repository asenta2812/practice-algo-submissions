func sortColors(nums []int) {
    countColors := [3]int{0,0,0} // red,white,blue

	for _, num := range nums {
		countColors[num]++
	}

	i:=0
	for color, count := range countColors {
		for j:=0; j < count; j++ {
			nums[i] = color
			i++
		}
	}
}
