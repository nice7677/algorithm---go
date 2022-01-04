package main

import "fmt"

func main() {
	array := []int{7, 5, 9, 0, 3, 1, 6, 2, 4, 8}
	fmt.Println(selectSort(array))
	/**
	[0 1 2 3 4 5 6 7 8 9]
	*/
}

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// 처음에 제일 작은거 설정
		minIndex := i
		for j := i; j < len(arr); j++ {
			// 전체 for loop 하면서 더 작은거 찾음
			if arr[minIndex] > arr[j] {
				// 찾으면 index 바꾸기
				minIndex = j
			}
		}
		// 값 스왑
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
