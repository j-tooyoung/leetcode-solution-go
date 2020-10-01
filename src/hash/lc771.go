package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	set := map[byte]bool{}
	res := 0
	for i := range J {
		set[J[i]] = true
	}
	for i := range S {
		if set[S[i]] {
			res++
		}
	}
	return res
}

func main() {

	J:= "Aa"
	S:= "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))


}
