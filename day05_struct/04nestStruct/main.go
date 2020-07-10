package main

import "fmt"

//结构体嵌套

type Address struct {
	Province string
	City     string
}

type Person struct {
	Name   string
	Age    int
	Gender string
	Address
}

func main() {
	var p1 = Person{
		Name:   "小天使",
		Age:    18,
		Gender: "男",
		Address: Address{
			"广东省", "广州市",
		},
	}

	//main.Person{Name:"小天使", Age:18, Gender:"男", Address:main.Address{Province:"广东省", City:"广州市"}}
	fmt.Printf("p1= %#v\n", p1)
	fmt.Println(p1.Address)          //{广东省 广州市}
	fmt.Println(p1.Address.Province) //广东省 通过嵌套内部的匿名结构体来访问其内部字段
	fmt.Println(p1.Province)         //广东省 直接通过p1访问嵌套结构体的内部字段
	fmt.Println(p1.City)             //广州市

}
