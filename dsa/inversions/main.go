package main

import (
	"dsa"
	"fmt"
	"math/rand"
)

func main() {
	a := dsa.RandIntSliceToN(rand.Intn(100))
	fmt.Println("Counting inversions for array: ", a)
	sorted, counted := sortAndCountInversions(a)
	fmt.Println("Inversions: ", counted)
	fmt.Println("Sorted: ", sorted)

	// Assignment
	inputs := dsa.ReadArrayFromTxt("inputs.txt")
	_, countedA := sortAndCountInversions(inputs)
	fmt.Println("Counted assignment: ", countedA)
}

func sortAndCountInversions(a []int) ([]int, int) {
	if len(a) <= 1 {
		return a, 0
	}

	m := len(a) / 2

	sortedLeft, leftInv := sortAndCountInversions(a[:m])
	sortedRight, rightInv := sortAndCountInversions(a[m:])
	sortedMerged, splitInv := mergedAndCountSplitInv(sortedLeft, sortedRight, len(a))
	return sortedMerged, leftInv + rightInv + splitInv
}

func mergedAndCountSplitInv(a, b []int, l int) ([]int, int) {
	var j, k, inv int
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
			inv += len(a) - j
		} else {
			m = append(m, a[j])
			j++
		}
	}
	return m, inv
}
