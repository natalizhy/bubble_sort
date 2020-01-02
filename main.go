package main

import (
	"fmt"
)

func main() {

	slice := []int{3, 4, 1, 2, 5, 7, -1, 0}
	fmt.Println("\n--- Unsorted --- \n\n", slice)

	order := false
	bubblesort(slice, order)
	fmt.Println("\n--- Sorted ---\n\n", slice)

	order = true
	bubblesort(slice, order)
	fmt.Println("\n--- Sorted2 ---\n\n", slice, "\n")


}

func bubblesort(items []int, order bool) {
	var (
		n = len(items)
	)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if !order {
				if items[i] > items[j] {
					items[j], items[i] = items[i], items[j]
				}
			} else {
				if items[j] > items[i] {
					items[i], items[j] = items[j], items[i]
				}
			}

		}
	}
}