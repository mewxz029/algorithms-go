package main

import (
	"fmt"

	"github.com/mewxz029/algorithms_go/algorithms"
)

func main() {
	fmt.Println("INSERT SORT")
	A := []int{5, 2, 4, 6, 1, 3}

	fmt.Println("Before sort: ", A)
	// TODO: Sub-array (sorted) start from right side
	// 5, 2, 4, 6, 1, 3
	// 5, 2, 4, 6, 1, 3
	// 5, 2, 4, 1, 3, 6
	// 5, 2, 1, 3, 4, 6
	// 5, 1, 2, 3, 4, 6
	// 1, 2, 3, 4, 5, 6
	algorithms.InsertSortDecreasing(A)
	fmt.Println("After sort: ", A)

	fmt.Println("LINEAR SEARCH")
	fmt.Println("Index of 4: ", algorithms.LinearSearch(A, 4))

	fmt.Println("ADD BINARY")
	A = []int{1, 0, 1}
	B := []int{1, 1, 1}

	fmt.Println("Add A + B = ", algorithms.AddBinary(A, B))
}
