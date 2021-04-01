package main

import "fmt"

/**
Go语言中的方法：
相当于Java中的成员变量，go语言中的方法可以给特定的类型添加特定的函数
*/

/**
鸟类结构体
*/
type Bird struct {
	name, color string // 名称，颜色
	age         int8   // 年龄
}

// 添加构造函数
func NewBird(name, color string, age int8) *Bird {
	return &Bird{
		name:  name,
		color: color,
		age:   age,
	}
}

// 为Bird类型添加一个“飞翔”方法
// b是接收者，这里指定了这个fly函数是Bird类型才能调用的函数
func (b Bird) fly() {
	fmt.Printf("我是%s,我在飞翔！！\n", b.name)
}

// 接收者可以分为指针接收者和值接收者
// 设置Bird年龄的方法
func (b *Bird) SetAge(newAge int8) {
	b.age = newAge
}

func (b Bird) SetAge1(newAge int8) {
	b.age = newAge
}

func main() {
	// 通过构造函数创建一个Bird实例
	bird1 := NewBird("鹦鹉", "绿色", int8(3))

	// bird1的类型和值: &main.Bird{name:"鹦鹉", color:"绿色", age:3}
	fmt.Printf("bird1的类型和值: %#v\n", bird1)

	// 调用飞翔方法
	(*bird1).fly() // 我是鹦鹉,我在飞翔！！
	// 语法糖：Go语言中，可以直接通过指针访问方法或者成员变量
	bird1.fly()

	// 改变年龄
	bird1.SetAge(int8(4))
	// bird1的年龄是：4
	fmt.Printf("bird1的年龄是：%d\n", bird1.age)

	// 通过值接收者的方法设置属性，设置不成功
	bird1.SetAge1(int8(5))
	// bird1的年龄是：4
	fmt.Printf("bird1的年龄是：%d\n", bird1.age)

	// 原因：接收者和函数参数一样，是值拷贝，如果传入的是结构体，结构体是值类型，调用的是一个新的结构体实例
	// 而不是原来那个

	// 使用指针型接收者有3种情况：
	// 1.需要方法中修改成员变量
	// 2.结构体成员变量很多，拷贝成本很大
	// 3.已经有其他方法使用了指针接收者
}
