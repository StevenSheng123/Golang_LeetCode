package main

//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	input := ""
//	for scanner.Scan() {
//		input += scanner.Text() + " "
//	}
//	data := strings.Fields(input)
//	index := 0
//	l, _ := strconv.Atoi(data[index])
//	index++
//	vec := make([]int, l)
//	for i := 0; i < l; i++ {
//		vec[i], _ = strconv.Atoi(data[index])
//		index++
//	}
//
//	//计算前缀和
//	p := make([]int, l)
//	preSum := 0
//	for i := 0; i < l; i++ {
//		preSum += vec[i]
//		p[i] = preSum
//	}
//	//处理查询
//	var results []int
//	for index < len(data) {
//		a, _ := strconv.Atoi(data[index])
//		fmt.Println("a:", a)
//
//		b, _ := strconv.Atoi(data[index+1])
//		fmt.Println("b:", b)
//		index += 2
//		if a == 0 {
//			results = append(results, p[b])
//		} else {
//			results = append(results, p[b]-p[a-1])
//		}
//
//	}
//	fmt.Println(p)
//	fmt.Println(results)
//	return
//
//}
