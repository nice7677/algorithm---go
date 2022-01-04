package main

import (
	"fmt"
	"sort"
)

func main() {
	N, K := 5, 3
	A := make([]int, N)
	A = []int{
		1, 2, 5, 4, 3,
	}
	B := make([]int, N)
	B = []int{
		5, 5, 6, 6, 5,
	}
	fmt.Println(두_배열의_원소_교체(A, B, K))
	/**
	26
	*/
}

func 두_배열의_원소_교체(A []int, B []int, K int) int {
	sum := 0
	sort.Ints(A) // [1,2,3,4,5]
	sort.Ints(B) // [5,5,5,6,6]
	for i := 0; i < K; i++ {
		// A는 앞에서부터 i씩 더하고, B는 뒤에서 부터 i씩 빼서 스왑
		A[i], B[len(B)-1-i] = B[len(B)-1-i], A[i]
	}
	for _, v := range A {
		sum += v
	}
	return sum
}
