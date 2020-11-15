package stack

import "fmt"

// 402. 移掉K位数字
//"1234567890" k=9
//"100204"     k =1
//"10"         k=1
func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	var list []byte
	cnt := 0
	list = append(list, num[0])
	flag := false
	for i := 1; i < len(num); i++ {
		//
		for !flag && len(list) > 0 && list[len(list)-1] > num[i] {
			list = append(list[:len(list)-1])
			//list = append(list[:len(list)-1], num[i])
			cnt++
			if cnt == k {
				flag = true
			}
		}
		list = append(list, num[i])
	}
	fmt.Println(cnt, k)

	i := 0
	for i < len(list) && list[i] == '0' {
		i++
	}
	if i >= len(list) {
		return "0"
	}
	// k -cnt
	// len -k +cnt
	var s string
	for ; i < len(list)-k-cnt; i++ {
		s = s + string(list[i])
	}
	return s
}
