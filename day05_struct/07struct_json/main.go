package main

import (
	"encoding/json"
	"fmt"
)

/**
结构体字段的可见性：

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）

下面的例子中，将会用json包中的函数，将班级学生数据序列化与反序列化，同时测试如果结构体字段用小写，json包
能否正常访问到
*/

type student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func newStudent(id int, name string) student {
	return student{
		Id:   id,
		Name: name,
	}
}

type class struct {
	//结构体Tag，格式：严格的键值对，多个Tag通过空格隔开
	ClassName string    `json:"class_name"`
	Students  []student `json:"students" db:"ss"`
}

func main() {
	//创建班级
	c1 := class{
		ClassName: "火箭101",
		Students:  make([]student, 0, 100),
	}
	//为班级添加学生
	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}

	//原始数据：
	fmt.Println("原始数据：")
	fmt.Printf("%#v\n", c1)
	fmt.Println("=========")

	//转json,注意需要传入指针
	data, err := json.Marshal(&c1)
	if err != nil {
		fmt.Println("转Json失败，error:", err)
		return
	}
	fmt.Println("转换的Json字符串")
	fmt.Printf("数据格式：%T\n", data)
	//当class的className和students两个字段都是首字母小写时，无法获得正确的JSON
	fmt.Printf("转为字符串输出：\n%s\n", data)

	//反序列化
	jsonStr := `{"class_name":"火箭102","students":[{"id":0,"name":"张三"},{"id":1,"name":"李四"},{"id":2,"name":"王五"}]}`
	var c2 class
	err1 := json.Unmarshal([]uint8(jsonStr), &c2)
	if err1 != nil {
		fmt.Println("反序列化失败：error:", err1)
		return
	}
	fmt.Println("反序列化：")
	fmt.Printf("%#v\n", c2)

}
