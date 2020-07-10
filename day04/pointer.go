package main

import "fmt"

func main() {
	n := 10
	fmt.Println(&n) //0xc000062080
	//把内存地址赋值给变量，这个变量就是一个指针
	p := &n
	//p的类型：*int,p的值=0xc000062080,根据p取值10
	fmt.Printf("p的类型：%T,p的值=%v,根据p取值%d\n", p, p, *p)
}
