package main

import (
	"dsa"
	"fmt"
	"slices"
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

	input := dsa.ReadArrayFromTxt("quicksort.txt")
	counter := &Counter{}

	sortedFirst := quickSortFirstElement(copySlice(input), counter)
	printSliceMeta(sortedFirst)
	fmt.Println("Num of comparisons - First elem: ", counter.Total())
	counter.Reset()

	sortedLast := quickSortLastElement(copySlice(input), counter)
	printSliceMeta(sortedLast)
	fmt.Println("Num of comparisons - Last elem: ", counter.Total())
	counter.Reset()

	sortedMedian := quickSortMedian(copySlice(input), counter)
	printSliceMeta(sortedMedian)
	fmt.Println("Num of comparisons - Median elem: ", counter.Total())

}

func printSliceMeta(s []int) {
	fmt.Println("Length: ", len(s), "Sorted: ", slices.IsSorted(s), "All Unique: ", len(s) == len(slices.Compact(s)))
}

type Counter struct {
	count int
}

func (c *Counter) Add(i int) {
	c.count += i
}

func (c *Counter) Total() int {
	return c.count
}

func (c *Counter) Reset() {
	c.count = 0
}

func copySlice(o []int) []int {
	new := make([]int, len(o))
	copy(new, o)
	return new
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

// Count number of comparisons, whenever there is a recursive call on
// a sub-array of lenght m add m-1 to the running total number
// of comparisons. Three scenarios:
// (1) Pivot is always the first element
func quickSortFirstElement(a []int, c *Counter) []int {

	if len(a) <= 1 {
		return a
	}

	// Increment comparisons counter
	c.Add(len(a) - 1)

	// Choose pivot
	p := a[0]

	// Partition
	j := 1
	for idx, i := range a[1:] {
		if i < p {
			firstUnpartitionedElem := a[j]
			a[j] = i
			a[idx+1] = firstUnpartitionedElem
			j++
		}
	}
	// Swap pivot with the boundary element (the rightmost element smaller than the pivot)
	a[0], a[j-1] = a[j-1], a[0]

	leftSorted := quickSortFirstElement(a[:j-1], c)
	righSorted := quickSortFirstElement(a[j:], c)
	return append(append(leftSorted, p), righSorted...)
}

// (2) Pivot is always the last element
func quickSortLastElement(a []int, c *Counter) []int {

	if len(a) <= 1 {
		return a
	}

	// Increment comparisons counter
	c.Add(len(a) - 1)

	// Choose pivot and swap with first element
	p := a[len(a)-1]
	a[0], a[len(a)-1] = a[len(a)-1], a[0]

	// Partition
	j := 1
	for idx, i := range a[1:] {
		if i < p {
			firstUnpartitionedElem := a[j]
			a[j] = i
			a[idx+1] = firstUnpartitionedElem
			j++
		}
	}
	// Swap pivot with the boundary element (the rightmost element smaller than the pivot)
	a[0], a[j-1] = a[j-1], a[0]

	leftSorted := quickSortLastElement(a[:j-1], c)
	righSorted := quickSortLastElement(a[j:], c)
	return append(append(leftSorted, p), righSorted...)
}

// (3) Pivot is median of first, middle and last element,
// if array is of even length, 2k, choose the k-th element as the middle one
func quickSortMedian(a []int, c *Counter) []int {
	if len(a) <= 1 {
		return a
	}

	// Increment comparisons counter
	c.Add(len(a) - 1)

	// Choose pivot and swap with first element
	p, idx := choosePivotMedian(a)
	a[0], a[idx] = a[idx], a[0]

	// Partition
	j := 1
	for idx, i := range a[1:] {
		if i < p {
			firstUnpartitionedElem := a[j]
			a[j] = i
			a[idx+1] = firstUnpartitionedElem
			j++
		}
	}

	// Swap pivot with the boundary element (the rightmost element smaller than the pivot)
	a[0], a[j-1] = a[j-1], a[0]

	leftSorted := quickSortMedian(a[:j-1], c)
	righSorted := quickSortMedian(a[j:], c)
	return append(append(leftSorted, p), righSorted...)
}

func choosePivotMedian(arr []int) (pivot, idx int) {
	length := len(arr)
	firstIdx := 0
	midIdx := (length - 1) / 2
	lastIdx := length - 1

	first := arr[firstIdx]
	mid := arr[midIdx]
	last := arr[lastIdx]

	// Median-of-three selection
	if (first >= mid && first <= last) || (first <= mid && first >= last) {
		return first, firstIdx
	}
	if (mid >= first && mid <= last) || (mid <= first && mid >= last) {
		return mid, midIdx
	}
	return last, lastIdx
}
