package main

import (
	"fmt"
)

func main() {
	// for i := range 1000 {
	// 	a := dsa.RandIntSlice(i)

	// 	naiveSorted := naiveQuickSort(a)
	// 	b := make([]int, len(a))
	// 	copy(b, a)
	// 	inplaceSorted := inPlaceQuickSort(b)

	// 	fmt.Println("Initial: ", a)
	// 	fmt.Println("Naive Sorted: ", naiveSorted)
	// 	fmt.Println("Naive Correct: ", slices.IsSorted(naiveSorted))
	// 	fmt.Println("InPlace Sorted: ", inplaceSorted)
	// 	fmt.Println("InPlace Correct: ", slices.IsSorted(inplaceSorted))
	// 	if !slices.IsSorted(inplaceSorted) {
	// 		panic("not sorted")
	// 	}
	// }
	a := []int{412, 652, 327, 95, 470, 205}
	b := inPlaceQuickSort(a)
	fmt.Println("Result: ", b)
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
	fmt.Println("initial", a)
	if len(a) <= 1 {
		return a
	}
	if len(a) == 2 {
		if a[0] < a[1] {
			return a
		}
		return []int{a[1], a[0]}
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
	boundaryElem := a[j]
	a[j] = p
	a[0] = boundaryElem

	leftSorted := inPlaceQuickSort(a[:j])
	righSorted := inPlaceQuickSort(a[j+1:])
	fmt.Println(leftSorted, righSorted)
	return append(append(leftSorted, p), righSorted...)
}
