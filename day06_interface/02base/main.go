package main

import "fmt"

/**
使用值接收者实现接口 VS 使用指针接收者实现接口的区别
 */

//接口：move
type IMover interface {
	move()
}

//Dog结构体，使用值接收者实现接口
type Dog struct {

}
//值接收者实现move
func (d Dog)move()  {
	fmt.Println("狗会跑~")
}

type Cat struct {

}
//使用指针接收者实现接口
func (c *Cat)move()  {
	fmt.Println("猫会跑~")
}

func main() {
	/**
	这里使用值接收者实现move时，可以发现值变量d1和Dog指针变量d2都可以赋值给m1
	这是因为Go底层有语法糖，自动取（*m1）
	 */
	var m1 IMover
	var d1 = Dog{}
	m1 = d1
	var d2 = &Dog{}
	m1 = d2
	m1.move()

	/**
	指针接收者实现接口
	 */
	var m2 IMover
	//var c1 = Cat{}
	var c2 = &Cat{}
	/**
	由于Cat实现接口的是*Cat类型，所以不能传入cat类型的c1
	cannot use c1 (type Cat) as type IMover in assignment:
	        Cat does not implement IMover (move method has pointer receiver)

	*/
	//m2 = c1
	m2 = c2
	m2.move()
}
