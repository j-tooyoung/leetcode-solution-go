package string

import "fmt"

//
func commonChars(A []string) []string {
	return nil
}

// 暴力法
func commonChars1(A []string) []string {
	var res []string
	if len(A) == 1 {
		for i := 0; i < len(A[0]); i++ {
			res =append(res, string(A[0][i]))
		}
	}
	common := A[0]
	for i := 1; i < len(A); i++ {
		common = commonChar(common,A[i])
		if len(common) < 1 {
			break
		}
	}
	for i := 0; i < len(common); i++ {
		res = append(res, string(common[i]))
	}
	return res
}

func commonChar(s1 string, s2 string) string {
	var ans string
	var cnt1 [26]int
	var cnt2 [26]int
	for i := 0; i < len(s1); i++ {
		cnt1[s1[i]-'a']++
	}
	for i := 0; i < len(s2); i++ {
		cnt2[s2[i]-'a']++
	}
	for i := 0; i < len(cnt1); i++ {
		if cnt1[i] > 0 && cnt2[i] > 0{
			cnt := min(cnt2[i],cnt1[i])
			for j := 0; j < cnt; j++ {
				ans =ans + string(i +'a')
			}
		}
	}
	fmt.Println(ans)
	return ans
}

func min(i int, j int) int {
	if i < j {
		return i
	}else {
		return j
	}
}
