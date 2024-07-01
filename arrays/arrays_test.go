package arrays

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraversing(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			15,
		},
		{},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, TraverseAdd(tc.input))
	}
}

func TestFindMax(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{10, 100, 300, 2},
			300,
		},
		{
			[]int{100, 500, 900, 2},
			900,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, FindMax(tc.input))
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
		{
			[]int{10, 20, 30, 40, 50},
			[]int{50, 40, 30, 20, 10},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, Reverse(tc.input))
	}
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		input1   []int
		input2   []int
		expected []int
	}{
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{10, 20, 30},
			[]int{40, 50, 60},
			[]int{10, 20, 30, 40, 50, 60},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, Merge(tc.input1, tc.input2))
	}
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		input    []int
		offset   int
		expected []int
		err      string
	}{
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			offset:   2,
			expected: []int{5, 6, 1, 2, 3, 4},
			err:      "Test Case 1 Failed: Basic rotation by 2 steps not handled correctly",
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			offset:   6,
			expected: []int{1, 2, 3, 4, 5, 6},
			err:      "Test Case 2 Failed: Rotation by length of array not handled correctly",
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			offset:   0,
			expected: []int{1, 2, 3, 4, 5, 6},
			err:      "Test Case 3 Failed: Rotation by 0 steps not handled correctly",
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			offset:   8,
			expected: []int{5, 6, 1, 2, 3, 4},
			err:      "Test Case 4 Failed: Rotation by more than length of array not handled correctly",
		},
		{
			input:    []int{},
			offset:   3,
			expected: []int{},
			err:      "Test Case 5 Failed: Empty array not handled correctly",
		},
		{
			input:    []int{1},
			offset:   5,
			expected: []int{1},
			err:      "Test Case 6 Failed: Single element array not handled correctly",
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			offset:   -1,
			expected: []int{2, 3, 4, 5, 6, 1},
			err:      "Test Case 7 Failed: Negative rotation not handled correctly",
		},
		{
			input:    []int{7, 7, 7, 7},
			offset:   3,
			expected: []int{7, 7, 7, 7},
			err:      "Test Case 9 Failed: Array with all same elements not handled correctly",
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			offset:   7,
			expected: []int{6, 1, 2, 3, 4, 5},
			err:      "Test Case 10 Failed: Rotation by array length plus one not handled correctly",
		},
	}

	for _, tc := range testCases {
		rotated := Rotate(tc.input, tc.offset)
		if !reflect.DeepEqual(rotated, tc.expected) {
			t.Logf("Expected %v, got %v", tc.expected, rotated)
			t.Fatal(tc.err)
		}
	}
}

func TestFindDuplicate(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 2},
			[]int{2},
		},
		{
			[]int{10, 20, 30, 40, 50, 10},
			[]int{10},
		},
		{
			[]int{10, 20, 30, 40, 50, 10, 20},
			[]int{10, 20},
		},
	}

	for _, tc := range testCases {
		duplicates := FindDuplicate(tc.input)
		if !reflect.DeepEqual(duplicates, tc.expected) {
			t.Fatalf("Expected %v, got %v", tc.expected, duplicates)
		}
	}
}

func TestFindMinAndMax(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 5},
		},
		{
			[]int{10, 20, 30, 40, 50},
			[]int{10, 50},
		},
		{
			[]int{10, 20, 30, 40, 50, 5},
			[]int{5, 50},
		},
	}

	for _, tc := range testCases {
		min, max := FindMinAndMax(tc.input)
		assert.Equal(t, tc.expected[0], min)
		assert.Equal(t, tc.expected[1], max)
	}
}

func TestTrappedRainWater(t *testing.T) {
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

	for _, tc := range testCases {
		fmt.Printf("t: %v\n", tc.input)
		assert.Equal(t, tc.expected, TrappedRainWater(tc.input))
	}
}
