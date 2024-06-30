package main

import (
	"fmt"
)

func main() {
	
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{3, 0, 0, 2, 0, 4},
			10,
		},
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0},
			6,
		},
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0, 1},
			6,
		},
		{
			[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			5,
		},
	}

	for _, test := range testCases {
		fmt.Println(test)
	}
}
