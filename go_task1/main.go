package main

import (
	"fmt"
	"sort"
)

func main() {
	//136. 只出现一次的数字
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
	// 9. 回文数
	x := 121
	fmt.Println(isPalindrome(x))
	//20.有效的括号
	s := "()[]{}"
	fmt.Println(isValid(s))
	// 14. 最长公共前缀
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	// 26. 删除有序数组中的重复项
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums2))
	//56. 合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
	// 66. 加一
	digits := []int{1, 9, 9}
	fmt.Println(plusOne(digits))
	//1. 两数之和
	nums1 := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums1, target))
}

// 136. 只出现一次的数字
// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
func singleNumber(nums []int) int {
	//最优解,异或运算
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
	//第一反应解法
	// numMap := make(map[int] int)
	// for _, val := range nums {
	//     _, ok := numMap[val]
	//     if ok {
	//         delete(numMap, val)
	//     } else {
	//         numMap[val] = val
	//     }
	// }
	// var onece int
	// for _, val := range numMap {
	//     onece = val
	//     break
	// }
	// return onece
}

// 9. 回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。
func isPalindrome(x int) bool {
	// 第一反应解法,将数字转换为字符串,然后反转字符串,比较反转后的字符串和原字符串是否相等
	// 反转字符串
	// str := strconv.Itoa(x)
	// runes := []rune(str)
	// for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	//     runes[i], runes[j] = runes[j], runes[i]
	// }
	// return string(runes) == str
	// 官方解法,反转一半数字
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	//判断字符串长度,为奇数直接返回false
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	shortest := strs[0]
	for _, s := range strs[1:] {
		if len(s) < len(shortest) {
			shortest = s
		}
	}

	for i := range shortest {
		for _, str := range strs {
			if str[i] != shortest[i] {
				return shortest[:i]
			}
		}
	}
	return shortest
}

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// 56. 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		if intervals[i][0] > last[1] {
			merged = append(merged, intervals[i])
		} else if intervals[i][1] > last[1] {
			last[1] = intervals[i][1]
		}
	}

	return merged
}

// 66. 加一
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
