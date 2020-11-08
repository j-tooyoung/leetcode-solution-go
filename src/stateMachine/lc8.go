package stateMachine

/**
https://blog.csdn.net/helen920318/article/details/105297687
https://leetcode-cn.com/problems/string-to-integer-atoi/
*/
/**
1 确定自动机中需要考虑的数据分类
答：数据有4类：数字，正负符号，空格和其他
2 确定自动机中数据的状态
答：自动机最后返回的结果为数字＊正负符号，若是遇到其他则返回0，空格不影响结果。
3 确定自动机的开始和结束
答：当遇到数字，或者正负符号的时候，自动机开始。
当数字溢出，或者遇到其他，或者得到最后的结果，自动机结束。
4 确定自动机开始后的流程

 - 遇到数字，［自动机开始］，进行计算，［溢出返回或计算结束返回，自动机结束］
 - 遇到正负符号，［自动机开始］，符号保留
 - 遇到不是数字，不是空格的，直接返回0
 - 自动机开始］的状态下，遇到非数字，［自动机结束］
 */
func myAtoi(s string) int {
	flag := 1
	res := 0
	begin := false

	for _, v := range s {
		if v >= '0' && v <= '9' {
			if res > 214748364 || res == 214748364 && (v > '7') {
				return 2147483647
			}
			if res < -214748364 || res == -214748364 && (v > '8') {
				return -2147483648
			}
			res = res * 10 + flag* int (v - '0')
			begin = true
		}else{
			if begin {	//开始转换后遇到非数字停止
				break
			}else if v == '-' {
				flag = -1
				begin = true
			}else if v == '+' {
				flag = 1
				begin = true
			}else if v != ' ' {
				return 0
			}
		}
	}
	return res
}
