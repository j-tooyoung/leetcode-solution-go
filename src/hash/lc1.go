package main

func twoSum(nums []int, target int)  []int {
	mp:= map[int]int{}
	//var mp = map[int]int{}
	//for i := 0; i < len(nums); i++ {
	//	if _, ok := mp[target-nums[i]]; ok {
	//		return []int{i, mp[target-nums[i]]}
	//	}else {
	//		mp[nums[i]] = i
	//	}
	//}
	for i, num := range nums {
		if index, ok := mp[target-num]; ok {
			return []int{index, i}
		}
		mp[num] = i
	}
	return nil
}

func main() {

}
