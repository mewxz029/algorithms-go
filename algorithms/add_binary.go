package algorithms

func AddBinary(a []int, b []int) []int {
	if len(a) != len(b) {
		return nil
	}

	c := make([]int, len(a)+1)

	carry := 0

	for i := len(a) - 1; i >= 0; i-- {
		sum := a[i] + b[i] + carry
		c[i+1] = sum % 2
		carry = sum / 2
	}

	c[0] = carry

	return c
}
