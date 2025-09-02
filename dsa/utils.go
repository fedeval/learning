package dsa

import (
	"math/rand"
)

func RandIntSlice(i int) []int {
	nums := make([]int, i)

	for j := 0; j < i; j++ {
		nums[j] = rand.Intn(1000)
	}
	return nums
}

func RandIntSliceToN(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	return nums
}
