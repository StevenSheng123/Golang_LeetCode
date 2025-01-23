package main

import (
	"fmt"
	"sort"
)

func Split(nums []int) ([]int, []int) {

	var nums1, nums2 []int
	for _, v := range nums {
		if v < 0 {
			nums1 = append(nums1, v)
		} else {
			nums2 = append(nums2, v)
		}
	}
	return nums1, nums2
}
func threeSum(nums []int) [][]int {
	//定义二元数组，保存返回值
	var ret [][]int
	//排序得到顺序的数字列表
	sort.Ints(nums)
	fmt.Println(nums)
	//将数组从0分隔开，nums1<0,nums2>=0
	nums1, nums2 := Split(nums)
	//遍历nums1,nums2的每两个元素，使其相加
	for i, il := 0, len(nums1); i < il; i++ {
		for j, jl := i+1, len(nums1); j < jl; j++ {
			//判断两个值是否相等
			if nums1[i] != nums1[j] {
				data := nums1[i] + nums1[j]
				for _, v := range nums2 {
					if data == 0-v {
						ret = append(ret, []int{nums1[i], nums1[j], v})
					}
				}
			}

		}
	}
	for i, il := 0, len(nums2); i < il; i++ {
		for j, jl := i+1, len(nums2); j < jl; j++ {
			//判断两个值是否相等
			if nums2[i] != nums2[j] {
				data := nums2[i] + nums2[j]
				for _, v := range nums1 {
					if data == 0-v && data != 0 {
						ret = append(ret, []int{nums2[i], nums2[j], v})
					}
				}
			}
		}
	}
	return ret
}

//func main() {
//	nums := []int{-1, 0, 1, 2, -1, -4}
//	fmt.Println(threeSum(nums))
//}
