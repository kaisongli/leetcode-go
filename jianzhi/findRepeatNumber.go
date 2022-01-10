package jianzhi

/**
剑指 Offer 03. 数组中重复的数字
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
*/
//重复首选哈希
func findRepeatNumber(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	m := make(map[int]int)
	for _, v := range nums {
		if m[v] > 0 {
			return v
		}
		m[v] += 1
	}
	return -1
}
