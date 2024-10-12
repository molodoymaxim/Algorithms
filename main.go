package main

import (
	"Algorithms/sorting"
	"fmt"
)

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	sorting.BubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}
