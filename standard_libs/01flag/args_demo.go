package main

import (
	"fmt"
	"os"
)

// flag包用来开发命令行参数

//1.os.Args []string类型，存储了用户输入的命令行参数
func main() {
	for i, arg := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, arg)
	}
}

/*
C:\Jalen\Programming\GoPath\src\HelloGo\standard_libs\01flag>args_demo a b c d
args[0]=args_demo
args[1]=a
args[2]=b
args[3]=c
args[4]=d
*/
