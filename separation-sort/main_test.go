package main

import (
	"reflect"
	"testing"
)

type separationTest struct {
	unordered []int
	ordered   []int
}

func TestSeparationSort(t *testing.T) {
	tests := []separationTest{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 3}, []int{1, 2, 3, 3, 4}},
		{[]int{10, 9, 8, 7, 6}, []int{6, 7, 8, 9, 10}},
		{[]int{300, 17, 22, 3, 7}, []int{3, 7, 17, 22, 300}},
		{[]int{300, 3417, 222, 3, 70}, []int{3, 70, 222, 300, 3417}},
	}

	for _, test := range tests {
		output := SeparationSort(test.unordered)
		if !reflect.DeepEqual(test.ordered, output) {
			t.Errorf("Want %v, Got %v", test.ordered, output)
		}
	}
}
