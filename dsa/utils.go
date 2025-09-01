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
