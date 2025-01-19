package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for x := range matrix {
		matrix[x] = make([]int, n)
	}
	i, j := 0, 0 //当前遍历到的横坐标、纵坐标信息
	num := 1     //左上角为1
	//if n%2 == 0 {
	//	//n是偶数，即最后一次遍历的正方形是2*2
	//	currentLen := n //当前遍历的正方形的边长，初始为n，逐圈-2
	//	flag := 0       //flag表示当前遍历方向，0代表向右，1代表向下，2代表向左，3代表向右
	//	for ; currentLen > 0; currentLen -= 2 {
	//		for edge := 0; edge < 4; edge++ {
	//			//edge表示每一圈要遍历四条边
	//			switch flag {
	//			case 0:
	//				for m := 0; m < currentLen-1; m++ {
	//					//current-1为该边要赋值的次数，初始为n-1次
	//					matrix[i][j] = num
	//					num++
	//					j++
	//				}
	//				flag = (flag + 1) % 4
	//
	//			case 1:
	//				for m := 0; m < currentLen-1; m++ {
	//					//current-1为该边要赋值的次数，初始为n-1次
	//					matrix[i][j] = num
	//					num++
	//					i++
	//				}
	//				flag = (flag + 1) % 4
	//			case 2:
	//				for m := 0; m < currentLen-1; m++ {
	//					//current-1为该边要赋值的次数，初始为n-1次
	//					matrix[i][j] = num
	//					num++
	//					j--
	//				}
	//				flag = (flag + 1) % 4
	//			case 3:
	//				for m := 0; m < currentLen-1; m++ {
	//					//current-1为该边要赋值的次数，初始为n-1次
	//					matrix[i][j] = num
	//					num++
	//					i--
	//				}
	//				flag = (flag + 1) % 4
	//			}
	//
	//		}
	//		i += 1
	//		j += 1
	//	}
	//
	//} else {
	//	//n是奇数，即最后一次遍历的正方形是1*1
	//	currentLen := n //当前遍历的正方形的边长，初始为n，逐圈-2
	//	flag := 0       //flag表示当前遍历方向，0代表向右，1代表向下，2代表向左，3代表向右
	//	for ; currentLen > 1; currentLen -= 2 {
	//		for edge := 0; edge < 4; edge++ {
	//			//edge表示每一圈要遍历四条边
	//			switch flag {
	//			case 0:
	//				for m := 0; m < currentLen-1; m++ {
	//					//current-1为该边要赋值的次数，初始为n-1次
	//					matrix[i][j] = num
	//					num++
	//					j++
	//				}
	//				flag = (flag + 1) % 4
	//
	//			case 1:
	//				for m := 0; m < currentLen-1; m++ {
	//					//current-1为该边要赋值的次数，初始为n-1次
	//					matrix[i][j] = num
	//					num++
	//					i++
	//				}
	//				flag = (flag + 1) % 4
	//			case 2:
	//				for m := 0; m < currentLen-1; m++ {
	//					//current-1为该边要赋值的次数，初始为n-1次
	//					matrix[i][j] = num
	//					num++
	//					j--
	//				}
	//				flag = (flag + 1) % 4
	//			case 3:
	//				for m := 0; m < currentLen-1; m++ {
	//					//current-1为该边要赋值的次数，初始为n-1次
	//					matrix[i][j] = num
	//					num++
	//					i--
	//				}
	//				flag = (flag + 1) % 4
	//			}
	//
	//		}
	//		i += 1
	//		j += 1
	//	}
	//	matrix[n/2][n/2] = n * n
	//}
	currentLen := n //当前遍历的正方形的边长，初始为n，逐圈-2
	flag := 0       //flag表示当前遍历方向，0代表向右，1代表向下，2代表向左，3代表向右
	for ; currentLen > 1; currentLen -= 2 {
		for edge := 0; edge < 4; edge++ {
			//edge表示每一圈要遍历四条边
			switch flag {
			case 0:
				for m := 0; m < currentLen-1; m++ {
					//current-1为该边要赋值的次数，初始为n-1次
					matrix[i][j] = num
					num++
					j++
				}
				flag = (flag + 1) % 4

			case 1:
				for m := 0; m < currentLen-1; m++ {
					//current-1为该边要赋值的次数，初始为n-1次
					matrix[i][j] = num
					num++
					i++
				}
				flag = (flag + 1) % 4
			case 2:
				for m := 0; m < currentLen-1; m++ {
					//current-1为该边要赋值的次数，初始为n-1次
					matrix[i][j] = num
					num++
					j--
				}
				flag = (flag + 1) % 4
			case 3:
				for m := 0; m < currentLen-1; m++ {
					//current-1为该边要赋值的次数，初始为n-1次
					matrix[i][j] = num
					num++
					i--
				}
				flag = (flag + 1) % 4
			}

		}
		i += 1
		j += 1
	}
	if n%2 == 1 {
		matrix[n/2][n/2] = n * n
	}
	return matrix
}
func main() {
	n := 4
	for _, val := range generateMatrix(n) {
		fmt.Println(val)
	}
	//fmt.Println(generateMatrix(n))
}
