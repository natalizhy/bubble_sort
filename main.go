package main

import (
	"fmt"
)

func main() {

	slice := []int{3, 4, 1, 2, 5, 7, -1, 0}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubblesort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
	bubblesort2(slice)
	fmt.Println("\n--- Sorted2 ---\n\n", slice, "\n")

}

func bubblesort(items []int) {
	var (
		n = len(items)
	)
	for i := 0; i < n-1; i++ {
		for j := i; j < n; j++ {
			if items[i] > items[j] {
				items[j], items[i] = items[i], items[j]
			}
		}
	}
}

func bubblesort2(items []int) {
	var (
		n = len(items)
	)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
}
