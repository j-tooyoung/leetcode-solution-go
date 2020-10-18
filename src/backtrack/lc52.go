package backtrack


func totalNQueens(n int) int {
	columns := make([]bool, n)		// 列
	dia1 := make([]bool, 2 *n -1)	// 左上到右下
	dia2 := make([]bool, 2 *n -1)
	var ans int
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			ans++
			return
		}
		for col, hasQueen := range columns {
			d1, d2 := row + n - 1- col, row +col
			if hasQueen || dia1[d1] || dia2[d2] {
				continue
			}
			columns[col] = true
			dia1[d1] = true
			dia2[d2] = true
			backtrack(row + 1)
			columns[col] = false
			dia1[d1] = false
			dia2[d2] = false
		}
	}
	backtrack(0)
	return ans
}
