package main

import "fmt"

func main() {
	array := []int{7, 5, 9, 0, 3, 1, 6, 2, 9, 1, 4, 8, 0, 5, 2}
	count := make([]int, len(array))
	count = countingSort(array, count)
	for i := range count {
		for j := 0; j < count[i]; j++ {
			fmt.Print(i, " ")
		}
	}
}

func countingSort(arr []int, count []int) []int {
	for _, item := range arr {
		count[item] += 1
	}
	return count
}
