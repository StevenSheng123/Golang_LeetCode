package main

import (
	"fmt"
)

//暴力双循环
//func minSubArrayLen(nums []int, target int) int {
//	count := []int{}
//	for i := range nums {
//		sum := 0
//		for j := i; j < len(nums); j++ {
//			sum += nums[j]
//			if sum >= target {
//				count = append(count, j-i+1)
//				break
//			}
//		}
//	}
//	sort.Ints(count)
//	if len(count) != 0 {
//		return count[0]
//	}
//	return 0
//}

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			result = min(subLength, result)
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}
