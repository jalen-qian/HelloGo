package main

import (
	"fmt"
	"io"
	"os"
)

/**
  小练习：借助io.Copy()函数实现一个拷贝文件的函数
*/
func CopyFile(srcName, dstName string) (n int64, err error) {
	// 以只读的方式打开源文件
	srcFile, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open srcFile failed,err:%v\n", err)
		return
	}
	defer srcFile.Close()
	// 以写的方式打开目标文件
	dstFile, err := os.OpenFile(dstName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open dstFile failed,err:%v\n", err)
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func main() {
	n, err := CopyFile("./aa.txt", "./cc.txt")
	if err != nil {
		fmt.Printf("copy failed,err:%v\n", err)
		return
	}
	fmt.Printf("copy succeed,%d bytes has been copied\n", n)
}
