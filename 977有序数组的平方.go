package main

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	i, j, k := 0, len(nums)-1, len(nums)-1
	for i < j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm < rm {
			ans[k] = rm
			j--
		} else {
			ans[k] = lm
			i++
		}
		k--
	}
	return ans
}

//func main() {
//	nums := []int{-4, -1, 0, 3, 10}
//	fmt.Println(sortedSquares(nums))
//}
