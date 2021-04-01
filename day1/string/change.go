package main

import "fmt"

func main() {
	// 这首诗写错了，需要将“皓月”改为“明月”
	s := "床前皓月光，疑似地上霜。举头望皓月，低头思故乡。"
	fmt.Println(&s) // 0xc00004e1c0
	// 强制类型转换为rune数组数组
	strArr := []rune(s)
	// 遍历，将皓改为明
	for i, c := range strArr {
		if c == '皓' {
			strArr[i] = '明'
		}
	}
	// 强制转换回字符串
	s = string(strArr) // 床前明月光，疑似地上霜。举头望明月，低头思故乡。
	// 打印改变后的字符串和地址
	fmt.Println(s)
	fmt.Println(&s) // 0xc00004e1c0
}
