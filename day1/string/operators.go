package main

import (
	"fmt"
	"strings"
)

func main() {
	//求长度
	s1 := "hello world"
	length := len(s1)
	fmt.Printf("s1的长度是%d\n", length) //s1的长度是11
	//len返回int类型
	fmt.Printf("%T\n", length) //int

	//拼接字符串
	s2 := "面朝大海"
	s3 := "春暖花开"
	fmt.Println(s2 + "," + s3) //面朝大海,春暖花开
	s4 := fmt.Sprintf("%s,%s", s2, s3)
	fmt.Println(s4) //面朝大海,春暖花开

	//分割字符串
	//strings.Split 可以将字符串用特定符号分割成字符串数组
	s5 := "床前明月光，疑似地上霜，举头望明月，低头思故乡"
	s6 := strings.Split(s5, "，")
	fmt.Printf("%v\n", s6) //[床前明月光 疑似地上霜 举头望明月 低头思故乡]
	fmt.Println(s6[2])     //举头望明月

	//判断是否包含
	//不包含
	if strings.Contains(s5, "s") {
		fmt.Println("包含")
	} else {
		fmt.Println("不包含")
	}
	//判断前缀、后缀
	if strings.HasPrefix(s5, "床前") {
		fmt.Println("包含前缀\"床前\"")
	}
	if strings.HasSuffix(s5, "故乡") {
		fmt.Println("包含后缀\"故乡\"")
	}

	//判断子串出现的位置
	//子串"明月"出现的位置是：6
	//因为一个字符是一个字节，一个rune占3个字节
	fmt.Printf("子串\"明月\"出现的位置是：%d\n", strings.Index(s5, "明月"))
	//子串"明月"最后出现的位置是：45
	fmt.Printf("子串\"明月\"最后出现的位置是：%d\n", strings.LastIndex(s5, "明月"))

	//join
	var arr = []string{"hello", "world"}
	arr2 := strings.Join(arr, " ")
	//将切片中的字符串用“空格”拼接
	fmt.Println(arr2) //hello world
}
