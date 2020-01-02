package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubblesort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
	bubblesort2(slice)
	fmt.Println("\n--- Sorted2 ---\n\n", slice, "\n")

}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(20) - rand.Intn(20)
	}
	return slice
}

func bubblesort(items []int) {
	var (
		n = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func bubblesort2(items []int) {
	var (
		n = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n; i++ {
			for j := n - 1; j > i; j-- {
				if items[j-1] > items[j] {
					items[j-1], items[j] = items[j], items[j-1]
					sorted = true
				}
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
