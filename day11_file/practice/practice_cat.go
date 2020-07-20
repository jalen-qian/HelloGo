package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// 练习：利用文件读取 和 flag包里的命令行参数工具，实现类似于linux的cat命令
// 实现效果: 编译后，输入cat aa.txt bb.txt 依次在终端中输出文件内容

/**
  输出文件内容
*/
func cat(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
	}
	defer file.Close()
	//创建一个writer，往终端输出
	writer := bufio.NewWriter(os.Stdout)
	for {
		var buf [1024]byte
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		_, err = writer.WriteString(string(buf[:n]))
		if err != nil {
			fmt.Printf("write to stdout failed, err:%v\n", err)
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Printf("flush failed,err:%v\n", err)
		return
	}

}
func main() {

	////cat("aa.txt")
	//for _, arg := range os.Args {
	//	//输出文件内容
	//	//cat(arg)
	//	fmt.Println(arg)
	//}
	flag.Parse() //解析命令行参数
	for i := 0; i < flag.NArg(); i++ {
		cat(flag.Arg(i))
	}
	//fmt.Println(flag.NArg())
}
