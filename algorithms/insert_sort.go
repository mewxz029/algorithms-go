package algorithms

func InsertSortDecreasing(A []int) {
	for i := len(A) - 2; i >= 0; i-- {
		key := A[i]
		j := i + 1
		for j < len(A) && A[j] < key {
			A[j-1] = A[j]
			j++
		}
		A[j-1] = key
	}
}
