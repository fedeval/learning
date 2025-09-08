package main

import (
	"dsa"
	"fmt"
	"slices"
)

func main() {
	a := dsa.RandIntSlice(1000)

	sorted := naiveQuickSort(a)

	fmt.Println("Initial: ", a)
	fmt.Println("Sorted: ", sorted)
	fmt.Println("Correct: ", slices.IsSorted(sorted))
}

func naiveQuickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	p := a[0]

	left, right := []int{}, []int{}
	for _, i := range a[1:] {
		if i < p {
			left = append(left, i)
		} else {
			right = append(right, i)
		}
	}

	leftSorted := naiveQuickSort(left)
	righSorted := naiveQuickSort(right)

	return append(append(leftSorted, p), righSorted...)
}
