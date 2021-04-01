package main

import "fmt"

/**
面试题：分析下面的代码能否通过编译
结论：不可以，因为Student结构体是通过“指针接收者”来实现的接口
所以people变量不能接收Student{}，只能接收指针 &Student{}
*/
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
