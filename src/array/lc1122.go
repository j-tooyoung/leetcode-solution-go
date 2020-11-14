package array

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	mp := make(map[int]int)
	//等价于
	//var mp map[int]int		//错误 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	//mp = make(map[int]interface{})

	for index, val := range arr2 {
		mp[val] = index
	}
	sort.Slice(arr1, func(i, j int) bool {
		index1, ok1 := mp[arr1[i]]
		index2, ok2 := mp[arr1[j]]
		if !ok1 && !ok2 {
			return arr1[i] < arr1[j]
		}else if ok1 && ok2 {
			return index1 < index2
		}else if ok1 && !ok2 {
			return true
		}else {
			return false
		}
	})
	return arr1

}
