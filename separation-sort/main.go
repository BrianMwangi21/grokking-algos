package main

import (
	"fmt"
)

func smallestIndex(list []int) int {
	smallest := list[0]
	smallest_index := 0

	for i := 1; i < len(list); i++ {
		if list[i] < smallest {
			smallest = list[i]
			smallest_index = i
		}
	}

	return smallest_index
}

func SeparationSort(list []int) []int {
	size := len(list)
	result := make([]int, size)

	for i := 0; i < size; i++ {
		smallest := smallestIndex(list)
		result[i] = list[smallest]
		list = append(list[:smallest], list[smallest+1:]...)
	}

	return result
}

func main() {
	fmt.Println("Hello, Separation Sort and I run at O(n^2)")
}
