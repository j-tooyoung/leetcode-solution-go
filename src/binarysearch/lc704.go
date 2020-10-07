package main

func search(nums []int, target int) int {
 	i,j:=0,len(nums)-1
	for i < j {
		mid := (i + j) >> 1
		if  nums[mid] < target{
			i = mid + 1
		}else {
			j=mid
		}
	}
	if nums[i] == target {
		return i
	} else {
		return -1
	}
}

func main() {
	
}
