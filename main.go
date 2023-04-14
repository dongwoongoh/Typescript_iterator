package main

import (
	"fmt"
)

func generateCombinations(arr []int, r int) [][]int {
	if r == 0 {
		return [][]int{[]int{}}
	}

	if len(arr) == 0 {
		return [][]int{}
	}

	head := arr[0]
	tail := arr[1:]

	withoutHead := generateCombinations(tail, r)
	withHead := generateCombinations(tail, r-1)

	for i := 0; i < len(withHead); i++ {
		withHead[i] = append([]int{head}, withHead[i]...)
	}

	return append(withHead, withoutHead...)
}

func main() {
	n := 6
	r := 4

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	combinations := generateCombinations(arr, r)
	result := len(combinations)

	fmt.Println(n, r, result)
}
