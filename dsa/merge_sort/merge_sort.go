package main

import (
	"dsa"
	"fmt"
)

// MergeSort
// 1 - Split the array in two halves
// 2 - Sort recursively
// 3 - Merge the results of the sorting
func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	l := len(a) / 2

	x := mergeSort(a[:l])
	y := mergeSort(a[l:])

	return merge(x, y, len(a))
}

// Merge two slices by iterating throught them
// in parallel and always selecting the smallest current
// item in the iteration. Assume the slices are sorted
// in increasing order.
func merge(a, b []int, l int) []int {

	var j, k int
	var m []int
	for range l {
		if j == len(a) {
			m = append(m, b[k:]...)
			break
		}

		if k == len(b) {
			m = append(m, a[j:]...)
			break
		}

		if a[j] > b[k] {
			m = append(m, b[k])
			k++
		} else {
			m = append(m, a[j])
			j++
		}
	}
	return m
}

func main() {
	a := dsa.RandIntSlice(10)
	fmt.Println("a: ", a)
	m := mergeSort(a)
	fmt.Println("m: ", m)
}
