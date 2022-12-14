package main

import (
	"fmt"
	"sort"
)

func BinarySearch(list []int, item int) int {
	sort.Ints(list)
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2 
		guess := list[mid]

		if guess == item {
			return mid
		}

		if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println("Hello, Binary Search and I run at O(log n)")
}
