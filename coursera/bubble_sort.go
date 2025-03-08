package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func BubbleSort(arr []int) {
	n := len(arr)
	for pass := 0; pass < n-1; pass++ {
		swapped := false
		for i := 0; i < n-pass-1; i++ {
			if arr[i] > arr[i+1] {
				Swap(arr, i)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	fmt.Println("Enter up to 10 integers separated by spaces:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	var numbers []int

	for _, part := range parts {
		if len(numbers) >= 10 {
			break
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Invalid input, please enter integers only.")
			return
		}
		numbers = append(numbers, num)
	}

	BubbleSort(numbers)

	fmt.Println("Sorted numbers:", numbers)
}
