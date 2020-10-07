package main

func mySqrt(x int) int {
	i, j := 0, x/2+1
	for i < j {
		mid := (i + j + 1) >> 1
		if mid*mid > x {
			j = mid - 1
		} else {
			i = mid
		}
	}
	return i

}

func main() {

}
