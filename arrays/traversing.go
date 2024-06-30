package main

import (
	"fmt"
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

func TrappedRainWater(heights []int) (total int) {
	left_index := 0
	right_index := 0

	for i := left_index; i < len(heights); i++ {
		if (heights[i] == 0) {
			continue
		}

		left_index = i
		break;
	}

	for i := 1; i < len(heights); i++ {
		if heights[i] >= heights[left_index] {
			right_index = i
			break
		}
	}

	blocks_between := right_index - left_index - 1
	min_height := min(heights[left_index], heights[right_index])

	sum_of_heights_between := func() int {
		var total int

		for i := left_index; i < right_index-1; i++ {
			total += heights[i]
		}

		return total - min_height
	}()



	fmt.Println(heights)
	fmt.Println("min height:", min_height)
	fmt.Println("sum of heights between:", sum_of_heights_between)
	fmt.Println()
	fmt.Println()
	total = (blocks_between*min_height) - sum_of_heights_between



	// we are at the next height that is equal to or bigger than left_index
	// TODO: EDGE CASES
	// 1. nothing between them
	// 2.

	return
}
