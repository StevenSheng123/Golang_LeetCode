package main

func removeElement(nums []int, target int) ([]int, int) {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != target {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return nums[:slowIndex], slowIndex
}

//func main() {
//	nums := []int{3, 2, 2, 3}
//	target := 2
//	fmt.Println(removeElement(nums, target))
//}
