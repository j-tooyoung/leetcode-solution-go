package array

func sortArrayByParityII(A []int) []int {
	i,j := 0, 1
	for j < len(A) && i < len(A) {
		for i < len(A) && A[i] % 2 == 0 {
			i += 2
		}
		for j < len(A) && A[j] % 2 == 1 {
			j += 2
		}
		if i < len(A) && j < len(A) {
			A[i], A[j] = A[j], A[i]
			i += 2
			j += 2
		}
	}
	return A


}


func sortArrayByParityII1(a []int) []int {
	for i, j := 0, 1; i < len(a); i += 2 {
		if a[i]%2 == 1 {
			for a[j]%2 == 1 {
				j += 2
			}
			a[i], a[j] = a[j], a[i]
		}
	}
	return a
}

