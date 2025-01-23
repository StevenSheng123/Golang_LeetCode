package main

func maxArea(height []int) int {
	res := 0
	i := 0
	j := len(height) - 1
	for i < j {
		max1 := (j - i) * min(height[i], height[j])
		res = max(res, max1)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}
	return res
}

//func main() {
//	height := []int{1, 1}
//	//height := []int{1,1}
//	fmt.Println(maxArea(height))
//}
