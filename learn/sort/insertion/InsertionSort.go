package main

import "fmt"

func main() {
	array := []int{7, 5, 9, 0, 3, 1, 6, 2, 4, 8}
	fmt.Println(insertionSort(array))
	/**
	[0 1 2 3 4 5 6 7 8 9]
	*/
}

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}
