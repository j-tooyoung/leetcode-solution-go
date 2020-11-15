package array

/**
我们希望下一个数比当前数大，这样才满足“下一个排列”的定义。因此只需要将后面的「大数」与前面的「小数」交换，就能得到一个更大的数。比如 123456，将 5 和 6 交换就能得到一个更大的数 123465。
我们还希望下一个数增加的幅度尽可能的小，这样才满足“下一个排列与当前排列紧邻“的要求。为了满足这个要求，我们需要：
*/
/**
1、从后向前查找第一个相邻升序的元素对 (i,j)，满足 A[i] < A[j]。此时 [j,end) 必然是降序
2、在 [j,end) 从后向前查找第一个满足 A[i] < A[k] 的 k。A[i]、A[k] 分别就是上文所说的「小数」、「大数」
3、将 A[i] 与 A[k] 交换
4、可以断定这时 [j,end) 必然是降序，逆置 [j,end)，使其升序
5、如果在步骤 1 找不到符合的相邻元素对，说明当前 [begin,end) 为一个降序顺序，则直接跳到步骤 4

*/
/**
先找出最大的索引 k 满足 nums[k] < nums[k+1]，如果不存在，就翻转整个数组；
再找出另一个最大索引 l 满足 nums[l] > nums[k]；
交换 nums[l] 和 nums[k]；
最后翻转 nums[k+1:]。

作者：powcai
链接：https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-powcai/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func nextPermutation(nums []int) {
	length := len(nums)
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	//
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	// fmt.Println(i, j,k)
	//
	if i < 0 {
		for i := 0; i < length/2; i++ {
			nums[length-1-i], nums[i] = nums[i], nums[length-1-i]
		}
		return
	}
	//
	for k >= j && nums[k] <= nums[i] {
		k--
	}

	//
	nums[i], nums[k] = nums[k], nums[i]
	//
	k = length - 1
	// fmt.Println(i, j,k)
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
	return
}
