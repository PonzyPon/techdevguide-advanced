package rainwater

import (
	"math"
)

func trap(height []int) int {
	var totalRain int
	length := len(height)
	leftMaxArray, rightMaxArray := getMaxArrays(height)
	for i, val := range height {
		min := int(math.Min(float64(leftMaxArray[i]), float64(rightMaxArray[length-i-1])))
		if min > val {
			totalRain += min - val
		}
	}
	return totalRain
}

func getMaxArrays(arr []int) ([]int, []int) {
	leftMax := 0.0
	leftMaxArray := []int{}
	rightMax := 0.0
	rightMaxArray := []int{}
	length := len(arr)
	for i := 0; i < len(arr); i++ {
		leftMax = math.Max(leftMax, float64(arr[i]))
		leftMaxArray = append(leftMaxArray, int(leftMax))
		rightMax = math.Max(rightMax, float64(arr[length-1-i]))
		rightMaxArray = append(rightMaxArray, int(rightMax))
	}
	return leftMaxArray, rightMaxArray
}
