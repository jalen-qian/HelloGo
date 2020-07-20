package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
文件操作,读文件
*/
func read() {
	file, err := os.Open("./aa.txt")
	//处理错误
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}
	defer file.Close() //记得处理完了要关闭文件
	var buf [127]byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			fmt.Print(string(buf[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read file failed,err:%v\n", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

/**
通过bufio来读文件
bufio封装了一层，里面设置了缓冲区
*/

func readByBufio() {
	file, err := os.Open("./aa.txt")
	//处理错误
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}
	defer file.Close() //记得处理完了要关闭文件
	//返回一个reader指针
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadString('\n') //根据换行符来读取
		if err == io.EOF { //End Of File
			fmt.Print(data)
			return
		}
		if err != nil {
			fmt.Printf("read by bufio failed, err:%v\n", err)
			return
		}
		fmt.Print(data)
	}
}

/**
  通过ioutil读取文件，进一步封装，不需要手动打开和关闭文件，也不需要缓冲区
  能直接将整个文件的所有内容读出来
  所以如果文件内容太大，不建议用这种方式，内存消耗会很大
*/
func readByIoutil() {
	data, err := ioutil.ReadFile("./aa.txt")
	if err != nil {
		fmt.Printf("read file by ioutil failed, err:%v\n", err)
	}
	fmt.Println(string(data))
}

/**
  write by file
  os.OpenFile()可以以指定的模式打开文件，从而实现文件的写入
  打开模式如下：
  os.O_WRONLY	只写
  os.O_CREATE	创建文件
  os.O_RDONLY	只读
  os.O_RDWR	    读写
  os.O_TRUNC	清空
  os.O_APPEND	追加
*/
func write() {
	//打开一个文件，只能写|如果没有自动创建|每次打开清空
	file, err := os.OpenFile("aaa.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}
	defer file.Close()
	n, err := file.WriteString("沙河小王子")
	if err != nil {
		fmt.Printf("write file failed,err:%v\n", err)
		return
	}
	fmt.Printf("写入了%d个字节\n", n)
}

func writeByBufio() {
	file, err := os.OpenFile("aa.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	n, err := writer.WriteString("\n新写入的一行")
	if err != nil {
		fmt.Printf("write file failed,err:%v\n", err)
		return
	}
	//bufio存在内置的缓冲区，所以写完需要入盘，一定要记得写这一行
	err = writer.Flush()
	if err != nil {
		fmt.Printf("flush to file failed,err:%v\n", err)
		return
	}
	fmt.Printf("写入了%d个字节\n", n)
}

func writeByIoutil() {
	err := ioutil.WriteFile("./aaa.txt", []byte("我是新的一行"), 0666)
	if err != nil {
		fmt.Printf("write by ioutil failed,err:%v\n", err)
	}
}

func main() {
	//read()
	//readByBufio()
	//readByIoutil()
	//write()
	//writeByBufio()
	writeByIoutil()
}
