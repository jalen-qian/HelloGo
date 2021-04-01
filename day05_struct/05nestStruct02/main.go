package main

import "fmt"

// 结构体嵌套字段冲突
// 当一个结构体嵌套的多个子结构体有同名字段时，访问字段必须加上子结构体的字段名称

type Address struct {
	Province   string
	City       string
	UpdateTime string
}

type Email struct {
	Addr       string
	UpdateTime string
}

type Person struct {
	Name   string
	Age    int
	Gender string
	Address
	Email
}

func main() {
	p := Person{
		Name:   "小天使",
		Age:    18,
		Gender: "男",
		Address: Address{
			"广东省",
			"广州市",
			"2020-07-10",
		},
		Email: Email{
			Addr:       "jalen@foxmail.com",
			UpdateTime: "2018-01-01",
		},
	}

	fmt.Printf("%#v\n", p)
	// fmt.Println(p.UpdateTime) //编译不通过：ambiguous selector p.UpdateTime
	fmt.Println(p.Address.UpdateTime)
	fmt.Println(p.Email.UpdateTime)
}
