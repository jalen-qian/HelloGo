package _1struct_base

import "fmt"

// 定义结构体
type Person struct {
	name   string
	age    int
	idCard string
}

// 相同类型的属性可以写在同一行
type Person1 struct {
	name, idCard string
	age, top     int
}

func main() {
	var p1 Person = Person{"张三", 15, "33422233244"}
	fmt.Printf("p1的类型：%T\n", p1) // p1的类型：main.Person

	var p2 Person
	p2.name = "李四"
	p2.age = 18
	p2.idCard = "2324343434343"
	fmt.Println(p2) //{李四 18 2324343434343}

	p3 := Person1{"王五", "23234322", 15, 173}

	// p3的类型和值：main.Person1{name:"王五", idCard:"23234322", age:15, top:173}
	fmt.Printf("p3的类型和值：%#v\n", p3)

	// 匿名结构体
	var student struct {
		name string
		age  int
	}
	student.name = "Jason"
	student.age = 18

	// type of student : struct { name string; age int }{name:"Jason", age:18}
	fmt.Printf("type of student : %#v\n", student)
}
