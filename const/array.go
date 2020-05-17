package main

import (
	"fmt"
)

func main() {
	//go定义数组必须指定特定的数组长度和类型
	var arr = [5]float32{1.02, 2.03, 3.04, 4.05, 5.06}
	fmt.Printf("%2f\n", arr[1])
	var arr1 [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		arr1[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("数组的第%d号元素是%d\n", j, arr1[j])
	}

	/*多维数组*/
	var arr2 = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Printf("第2行3列的数字是%d\n", arr2[1][2])

	//向函数中传入数组
	var arr3 = [...]int{1, 3, 4, 5, 3, 3, 45, 66, 7, 442, 3, 4, 5667, 7, 8, 6, 5, 4, 43, 3, 3, 2, 43, 4, 5, 65, 6, 6, 54, 34, 3, 3, 32, 2, 2, 4444, 2, 23, 3, 3, 43, 4, 4}
	getAverage(arr3)
}

/**
* 测试向函数中传入数组，并计算平均值
 */
func getAverage(arr [43]int) {
	var avg float32
	var size, sum = len(arr), 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)

	fmt.Printf("平均值为：%f\n", avg)
}
