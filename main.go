package main

import (
	"fmt"
)

func generateCombinations(arr []uint64, r uint64) [][]uint64 {
	if r == 0 {
		return [][]uint64{[]uint64{}}
	}

	if len(arr) == 0 {
		return [][]uint64{}
	}

	head := arr[0]
	tail := arr[1:]

	withoutHead := generateCombinations(tail, r)
	withHead := generateCombinations(tail, r-1)

	for i := 0; i < len(withHead); i++ {
		withHead[i] = append([]uint64{head}, withHead[i]...)
	}

	return append(withHead, withoutHead...)
}

func main() {
	n := 8
	r := 4

	arr := make([]uint64, n)
	for i := 0; i < n; i++ {
		arr[i] = uint64(i + 1)
	}

	combinations := generateCombinations(arr, uint64(r))
	result := len(combinations)

	fmt.Println(n, r, result)
}
