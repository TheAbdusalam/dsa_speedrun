package arrays

import (
	"math"
)

func TraverseAdd(n []int) int {
	var total int
	for _, num := range n {
		total += num
	}

	return total
}

func FindMax(n []int) int {
	var max int

	for _, num := range n {
		if num > max {
			max = num
		}
	}

	return max
}

func Reverse(n []int) []int {
	var reversed []int

	for i := len(n) - 1; i >= 0; i-- {
		reversed = append(reversed, n[i])
	}

	return reversed
}

func Merge(arr1 []int, arr2 []int) []int {
	return append(arr1, arr2...)
}

func Rotate(arr []int, offset int) []int {

	if offset == 0 || len(arr) == 0 {
		return arr
	}

	if offset < 0 {
		offset = len(arr) + offset
	}

	offset = offset % len(arr)
	return append(arr[len(arr)-offset:], arr[:len(arr)-offset]...)

}

func FindDuplicate(arr []int) (result []int) {
	seen := make(map[int]bool)

	for _, num := range arr {
		if seen[num] {
			result = append(result, num)
		}

		seen[num] = true
	}

	return
}

func FindMinAndMax(arr []int) (min, max int) {
	if len(arr) == 0 {
		return
	}

	min = arr[0]
	max = arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	return
}

func TrappedRainWater(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	leftMax := make([]int, len(heights))
	rightMax := make([]int, len(heights))

	leftMax[0] = heights[0]
	for i := 1; i < len(heights); i++ {
		leftMax[i] = int(math.Max(float64(leftMax[i-1]), float64(heights[i])))
	}

	rightMax[len(heights)-1] = heights[len(heights)-1]
	for i := len(heights) - 2; i >= 0; i-- {
		rightMax[i] = int(math.Max(float64(rightMax[i+1]), float64(heights[i])))
	}

	var total int
	for i := 0; i < len(heights); i++ {
		total += int(math.Min(float64(leftMax[i]), float64(rightMax[i]))) - heights[i]
	}

	return total
}
