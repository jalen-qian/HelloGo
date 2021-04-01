package _1struct_base

import (
	"fmt"
	"unsafe"
)

type Person2 struct {
	name string
	age  int
	city string
}

func main() {
	// 结构体初始化
	// 1.当一个结构体没有初始化时，对应结构体成员变量都是对应类型的零值（也就是说此时也分配了内存，成员变量默认零值）
	var p1 Person2
	// p1的类型和值：main.Person2{name:"", age:0, city:""}
	fmt.Printf("p1的类型和值：%#v\n", p1)

	// 测试p1是否占用内存空间
	fmt.Println(unsafe.Sizeof(p1)) // 40,说明分配了空间

	// 2.可以通过键值对来初始化，如果某些属性不写，默认零值（这里的city属性）
	p2 := Person2{name: "张三", age: 18}
	// p2的类型和值：main.Person2{name:"张三", age:18, city:""}
	fmt.Printf("p2的类型和值：%#v\n", p2)

	// 也可以对结构体的指针进行键值对初始化
	p3 := &Person2{name: "李四", age: 19, city: "重庆"}
	// p3的类型和值：&main.Person2{name:"李四", age:19, city:"重庆"}
	fmt.Printf("p3的类型和值：%#v\n", p3)

	// 3.使用值列表初始化
	p4 := Person2{
		"王五",
		18,
		"重庆",
	}
	// p4=main.Person2{name:"王五", age:18, city:"重庆"}
	fmt.Printf("p4=%#v\n", p4)

	/** 空结构体 */
	//空结构体是不占用内存空间的

	var p5 struct{}
	// p5= struct {}{}
	fmt.Printf("p5= %#v\n", p5)
	// 测试p5是否占用内存空间
	// p5占用内存0
	fmt.Printf("p5占用内存%d\n", unsafe.Sizeof(p5))
}
