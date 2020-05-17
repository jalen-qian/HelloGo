package main

import (
	"fmt"
)

//定义结构体
type People struct {
	name string
	age  int
	job  string
}

func main() {
	var popople1 People = People{"张三", 25, "教师"}
	people2 := People{"李四", 45, "医生"}
	fmt.Println(popople1)
	fmt.Println(people2)

	fmt.Println("==================================")
	//局部定义结构体
	type Student struct {
		studentId int      //学生ID
		name      string   //学生名称
		age       int      //年龄
		grade     int      //年级
		classes   []string //主修课程
	}

	tom := Student{100001, "小明", 12, 3, []string{"语文", "数学", "英语"}}
	fmt.Println(tom)

	fmt.Println("==================================\n")

	//单个属性赋值
	var jerry Student
	jerry.studentId = 100002
	jerry.name = "jerry"
	jerry.age = 16
	jerry.grade = 4
	jerry.classes = []string{"语文", "数学", "英语", "科学"}
	fmt.Println(jerry)

	fmt.Println("==================================\n")
	fmt.Println("结构体指针\n")

	var mary Student = Student{100003, "mary", 17, 9, []string{"语文", "数学", "英语", "物理", "化学", "生物"}}
	fmt.Println(mary)
	var p *Student = &mary
	//输出 结构体指针，指向mary，地址为：&{186a3 6d617279 11 9 [e8afade69687 e695b0e5ada6 e88bb1e8afad e789a9e79086 e58c96e5ada6 e7949fe789a9]}
	fmt.Printf("结构体指针，指向mary，地址为：%x\n", p)
	//通过结构体指针输出mary的年龄
	fmt.Println(p.age) //17

}
