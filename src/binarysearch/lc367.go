package main

func isPerfectSquare(num int) bool {
	i, j := 0, num/2+1
	for i < j {
		mid := (i + j) >> 1
		//multi := mid * mid
		if mid*mid < num {
			i = mid + 1
		} else {
			j = mid
		}
	}
	if j*j == num {
		return true
	} else {
		return false
	}
}

func main() {

}
