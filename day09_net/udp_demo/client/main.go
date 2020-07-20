package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//udp客户端
func main() {
	//拨号，通过udp
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("dial udp failed,err:%v\n", err)
	}
	defer udpConn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		//发送数据到服务端
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from os.stdin failed,err:%v\n", err)
			break
		}
		if input == "Q" {
			return
		}
		n, err := udpConn.Write([]byte(input))
		if err != nil {
			fmt.Printf("send to server failed,err:%v\n", err)
			break
		}
		//获取服务端的响应
		var buff [1024]byte
		n, err = udpConn.Read(buff[:])
		if err != nil {
			fmt.Printf("read from server failed,err:%v\n", err)
			break
		}
		fmt.Printf("response from server:%s\n", string(buff[:n]))
	}
}
