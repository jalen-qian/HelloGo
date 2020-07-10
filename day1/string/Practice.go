package main

import (
	"fmt"
	"unicode"
)

//判断是否字符是汉字
func isHan(char rune) bool {
	return unicode.Is(unicode.Han, char)
}

func main() {
	/*
		练习题：编写程序，统计字符串 "hello沙河小王子なつめ"中汉字的个数
	*/
	str := "hello沙河小王子なつめ"
	var hanNums int
	for _, c := range str {
		if isHan(c) {
			hanNums++
		}
	}
	fmt.Printf("汉字个数是%d\n", hanNums) //汉字个数是5

}
