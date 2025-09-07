package dsa

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

func ReadArrayFromTxt(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Prepare a slice to hold the numbers
	var numbers []int

	// Read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line) // convert string to int
		if err != nil {
			fmt.Println("Skipping invalid line:", line)
			continue
		}
		numbers = append(numbers, num)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return numbers
}
