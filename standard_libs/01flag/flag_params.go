package main

import (
	"flag"
	"fmt"
	"time"
)

/**
flag.Type可以定义命令行参数的名称、默认值、说明，并且返回一个对应类型的指针

执行结果如下：
C:\Jalen\Programming\GoPath\src\HelloGo\standard_libs\01flag>flag_params -name "王五" -age=19 -married=false -delay 1m20s
王五 19 false 1m20s

*/
func flagType() {
	// 通过flag.Type来定义参数
	// flag.Type(flag名, 默认值, 帮助信息)*Type 返回的是对应类型的指针
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("delay", 0, "延迟的时间间隔")

	// 通过flag.Parse()解析命令行参数
	flag.Parse()
	fmt.Println(*name, *age, *married, *delay)
}

/**
  flag.TypeVar()定义命令行参数
*/
func flagTypeVar() {
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "delay", 0, "延迟的时间间隔")
	flag.Parse()
	// 打印
	fmt.Println(name, age, married, delay)
	// 打印flag的其他参数
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())  // 返回命令行后的其他参数个数 0
	fmt.Println(flag.NFlag()) // 返回使用的参数个数  4
}

func main() {
	// flagType()
	flagTypeVar()
}
