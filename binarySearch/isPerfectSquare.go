package binarySearch
//给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
//
// 说明：不要使用任何内置的库函数，如 sqrt。
//
// 示例 1：
//
// 输入：16
//输出：True
//
// 示例 2：
//
// 输入：14
//输出：False
//
// Related Topics 数学 二分查找
// 👍 168 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	if num < 2{
		return true
	}
	left, right := 2, num/2
	for left <= right{
		mid := left + (right - left) >> 1
		if mid * mid == num{
			return true
		}else if mid * mid < num{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)

