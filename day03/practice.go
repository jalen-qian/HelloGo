package main

import "fmt"

/*
数组练习题：
1.求数组[1, 3, 5, 7, 8]所有元素的和
2.找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
假定数组中的元素只会出现一次
*/

func main() {
	//第一题
	arr1 := [...]int{1, 3, 5, 7, 8}
	result := 0
	for _, v := range arr1 {
		result += v
	}
	fmt.Printf("%v求和之后的结果是：%d\n", arr1, result) //[1 3 5 7 8]求和之后的结果是：24

	//第二题
	findIndex([10]int{1, 3, 5, 7, 8, 6, 2, 4, 9, 10}, 10)
	fmt.Println()

	findIndex1([10]int{1, 3, 5, 7, 8, 6, 2, 4, 9, 10}, 10)
}

/**
找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
解法一，双层遍历，时间复杂度O(n^2),空间复杂度O(1)
*/
func findIndex(arr [10]int, target int) {
	for i, v := range arr {
		for j := i + 1; j < len(arr); j++ {
			if v+arr[j] == target {
				fmt.Printf("找到下标(%d,%d)的结果\n", i, j)
			}
		}
	}
}

/**
解法二：
用一个map存储对应的值=>下标，然后只需遍历一次数组即可
时间复杂度O(n) 空间复杂度O(n)
*/
func findIndex1(arr [10]int, target int) {
	myMap := map[int]int{}
	for i, v := range arr {
		if j, isOk := myMap[target-v]; isOk {
			fmt.Printf("找到下标(%d,%d)的结果\n", j, i)
		}
		myMap[v] = i
	}
}
