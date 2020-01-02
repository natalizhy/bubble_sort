package main

import (
	"fmt"
)

func main() {

	slice := []int{3, 4, 1, 2, 5, 7, -1, 0}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubblesort(slice)

}

func bubblesort(items []int) {
	var (
		n = len(items)
	)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if items[i] > items[j] {
				items[j], items[i] = items[i], items[j]
			}
		}
	}
	fmt.Println("\n--- Sorted ---\n\n", items)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if items[j] > items[i] {
				items[i], items[j] = items[j], items[i]
			}
		}
	}
	fmt.Println("\n--- Sorted2 ---\n\n", items, "\n")

}
