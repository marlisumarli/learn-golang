package main

import (
	"fmt"
	"sort"
	"sync"
)

func merge(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)
	return result
}

func main() {
	var n int
	fmt.Print("Enter the number of integers: ")
	fmt.Scan(&n)

	numbers := make([]int, n)
	fmt.Println("Enter the numbers:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	partSize := (n + 3) / 4
	subarrays := make([][]int, 4)

	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if end > n {
			end = n
		}
		subarrays[i] = numbers[start:end]
	}

	var wg sync.WaitGroup
	wg.Add(4)

	for i := 0; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("Sorting Subarray:", subarrays[i])
			sort.Ints(subarrays[i])
			fmt.Println("Sorted Subarray:", subarrays[i])
		}(i)
	}

	wg.Wait()

	sortedArray := merge(merge(subarrays[0], subarrays[1]), merge(subarrays[2], subarrays[3]))

	fmt.Println("Final Sorted Array:", sortedArray)
}
