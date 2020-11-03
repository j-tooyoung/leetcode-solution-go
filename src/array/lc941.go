package array

func validMountainArray(A []int) bool {
	length := len(A)
	i, j := 1, length -1
	if length < 3 {
		return false
	}

	for i < length && A[i -1] < A[i] {
		i++
	}
	for j > 0 && A[j -1] > A[j] {
		j--
	}
	// fmt.Println(i, j)
	if j == length - 1 || i == 1 {
		return false
	}
	return i - j == 1

}
