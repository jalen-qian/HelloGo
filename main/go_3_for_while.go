package main

import "fmt"

/**
 * 计算一个范围内的质数并用切片返回
 * 时间复杂度 O(n^2)
 */
func findPrimeNum(scope int) []int {
	var a []int
	var isPrime bool
	//质数：大于1的自然数中，只能被1和自身整除的数是质数
	index := 0 //统计总共执行的次数，并打印
	for i := 2; i <= scope; i++ {
		isPrime = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				index++
				break
			} else {
				index++
			}
		}
		if isPrime {
			a = append(a, i)
		}
	}
	fmt.Printf("使用方法1总共执行了%d次\n", index)
	return a
}

/**
 * 更优解法，从质数定义可知，如果i不是质数，则存在整数a满足j * a = i，
 * 其中 j < a < i, 那么范围(a,i)之间的数就不需要遍历了
 * 比如寻找95这个数是否为质数时，i * a = 95, 因为i最小是2，所以a不可能超过 95/2 = 47
 * 也就是说 48,49...94这些数字都不用遍历了，减少了遍历的次数
 */
func findPrimeNum2(scope int) []int {
	var a []int
	index := 0 //统计总共执行的次数，并打印
	//质数：大于1的自然数中，只能被1和自身整除的数是质数
	var i, j int
	for i = 2; i <= scope; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				index++
				break
			} else {
				index++
			}
		}
		if j > (i / j) {
			a = append(a, i)
		}
	}
	fmt.Printf("使用方法2总共执行了%d次\n", index)
	return a
}

func printSlice(slice []int) {
	fmt.Printf("%v\n", slice)
}

/**
 * 使用
 */

func main() {
	//for 循环
	/*for true {
		fmt.Println("这是一个无限循环")
	}*/

	for i := 0; i < 10; i++ {
		fmt.Printf("这是%d次循环\n", i)
	}

	//题目：通过嵌套for循环输出一定整数范围内的质数
	a := findPrimeNum(1000)
	//输出结果
	printSlice(a)
	b := findPrimeNum2(1000)
	printSlice(b)

	//也可以不在for里面初始化值，或者省略掉递增（减）
	c := 0
	for ; c < 10; c++ {
		fmt.Printf("c=%d ", c)
	}
	c = 1
	for c < 100 {
		c += c
	}
	fmt.Println("")
	fmt.Printf("c=%d\n", c)

	//条件也不写时，进入无限循环
	c = 1
	//for {
	//	c++
	//}
	//fmt.Printf("c=%d\n",c)//这一行不会执行到，因为上面是无限循环

	/** For-Each Range循环，这种循环可以对字符串、数组、切片等进行迭代输出 */
	str := []string{"hello", "world"}
	for i, s := range str {
		fmt.Println(i, s) //0 hello  1 world
	}
	var arr = [5]int{1, 2, 3, 4, 5}
	for i, s := range arr {
		fmt.Println(i, s)
	}

	//break语句
	//break用于跳出循环，或者switch中用于执行一条case后跳出
	//这里讲一个特殊的，GoLang中的break可以通过label跳到多重循环中的任意一重

	for i1 := 1; i1 <= 3; i1++ {
		fmt.Printf("i1=%d\n", i1)
	la:
		for i2 := 11; i2 <= 33; i2 += 11 {
			fmt.Printf("i2=%d\n", i2)
			for i3 := 111; i3 <= 333; i3 += 111 {
				fmt.Printf("i3=%d\n", i3)
				break la
			}
		}
	}
	/*
		i1=1
		i2=11
		i3=111
		i1=2
		i2=11
		i3=111
		i1=3
		i2=11
		i3=111
		从结果可以看出，当第一次输出i3=111之后，由于使用了标签la,跳过了内层的两层循环，直接执行最外层的第二次循环
		再次执行到最内层的输出i3=111之后，又一次跳到最外层的第三次循环，所以总共输出了9行
	*/

}
