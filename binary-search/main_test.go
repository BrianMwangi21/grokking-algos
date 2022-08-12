package main

import (
	"testing"
)

type binaryTest struct {
	list []int
	item int
	expected int
}

func TestBinarySearch(t *testing.T) {
	tests := []binaryTest{
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 7, -1},
		{[]int{6, 7, 8, 9, 10}, 8, 2},
		{[]int{300, 17, 22, 3, 7}, 300, 4},
		{[]int{300, 3417, 222, 3, 70}, 121, -1},
	}

	for _, test := range tests {
		output := BinarySearch(test.list, test.item)
		if output != test.expected {
			t.Errorf("Want %v, Got %v", test.expected, output)
		}
	}
}
