
func solution (height []int) int {
	left := 0
	right := len(height) - 1
	currentMax := 0

	for {
		if left = right {
			break
		}

		width := right - left
		height := math.min(height[left], height[right])
		currentVol = width * height

		if currentVol > currentMax {
			currentMax = currentVol
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return currentMax
}