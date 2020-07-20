package main

import (
	"fmt"
	"net"
)

// udp服务器
func main() {
	// 1.监听udp服务器的ipv4127.0.0.1,3000端口
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("listen udp failed,err:%v\n", err)
		return
	}
	defer listen.Close() //结束之前关闭udp通信
	for {
		//2.不断读取客户端发送来的数据
		var buff [1024]byte
		n, addr, err := listen.ReadFromUDP(buff[:])
		if err != nil {
			fmt.Printf("read from udp addr %v failed,err:%v\n", addr, err)
			break
		}
		//打印读到的数据
		data := string(buff[:n])
		fmt.Printf("收到从客户端%v发送的数据：%s\n", addr, data)
		//将数据发送回客户端
		n, err = listen.WriteToUDP([]byte(data), addr)
		if err != nil {
			fmt.Printf("send to client failed,err:%v\n", err)
			break
		}
	}

}
