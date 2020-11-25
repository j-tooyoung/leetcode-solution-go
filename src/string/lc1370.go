package string

func sortString(s string) string {
	cnt:= make([]int, 26)
	for _, val := range s {
		cnt[val-'a']++
	}
	//sort.Ints(cnt)
	//var res string
	ans := make([]byte,0,len(s))
	for true {
		flag1 := false
		flag2 := false

		for i := 0; i < 26; i++ {
			if cnt[i] > 0 {
				ans = append(ans,byte('a'+i))
				//res = res + string('a'+ i)
				cnt[i]--
				flag1 = true
			}
		}
		for i := 25; i >=0 ; i-- {
			if cnt[i] > 0 {
				ans = append(ans,byte('a'+i))
				//res = res + string('a'+ i)
				cnt[i]--
				flag2 = true
			}
		}
		if !flag1 && !flag2 {
			break
		}
	}
	return string(ans)
}
