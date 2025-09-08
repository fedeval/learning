package main

import (
	"dsa"
	"fmt"
	"slices"
)

func main() {
	for i := range 1000 {
		a := dsa.RandIntSlice(i)

		naiveSorted := naiveQuickSort(a)
		b := make([]int, len(a))
		copy(b, a)
		inplaceSorted := inPlaceQuickSort(b)

		fmt.Println("Initial: ", a)
		fmt.Println("Naive Sorted: ", naiveSorted)
		fmt.Println("Naive Correct: ", slices.IsSorted(naiveSorted))
		fmt.Println("InPlace Sorted: ", inplaceSorted)
		fmt.Println("InPlace Correct: ", slices.IsSorted(inplaceSorted))
		if !slices.IsSorted(inplaceSorted) {
			panic("not sorted")
		}
	}
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

func inPlaceQuickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	p := a[0]

	j := 1
	for idx, i := range a[1:] {
		if i < p {
			firstUnpartitionedElem := a[j]
			a[j] = i
			a[idx+1] = firstUnpartitionedElem
			j++
		}
	}
	boundaryElem := a[j-1]
	a[j-1] = p
	a[0] = boundaryElem

	leftSorted := inPlaceQuickSort(a[:j-1])
	righSorted := inPlaceQuickSort(a[j:])
	return append(append(leftSorted, p), righSorted...)
}
