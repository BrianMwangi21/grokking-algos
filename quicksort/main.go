package main 

import "fmt"

func Quicksort(list []int) []int {
  size := len(list)

  if size < 2 {
    return list 
  }else {
    pivot := list[0]
    var less, greater []int

    for _, num := range list[1:] {
      if num < pivot {
        less = append(less, num)
      } else {
        greater = append(greater, num)
      }
    }

    less = append(Quicksort(less), pivot)
    greater = Quicksort(greater)
    return append(less, greater...)
  }
}

func main() {
	fmt.Println("Hello, Quicksort and I run at O(n log n)")
}
