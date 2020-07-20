package main

import (
	"fmt"
)

/**
一个接口的方法，不一定需要在一个类型中实现，可以在这个类型中嵌入其他类型或者结构体来实现
*/

//例子：洗衣机有洗和甩干功能，甩干是甩干机的功能，可以在洗衣机中嵌套甩干机来实现洗衣服接口

type washingMachine interface {
	wash() //洗衣服
	dry()  //甩干
}

//甩干机
type drier struct{}

//甩干机实现了甩干功能
func (d drier) dry() {
	fmt.Println("我在甩干衣服")
}

//海尔洗衣机
type haier struct {
	drier //海尔洗衣机内嵌套甩干机
}

func (h haier) wash() {
	fmt.Println("海尔洗衣机在洗衣服")
}

func main() {
	//能正常赋值，说明haier洗衣机实现了washingMachine接口
	var washingM washingMachine = haier{}
	//海尔洗衣机在洗衣服
	//我在甩干衣服
	washingM.wash()
	washingM.dry()
}
