package main

import "fmt"

/*
分金币：
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	left := coins
	// 匿名函数，计算名字对应的coin数量
	calcCoins := func(name string) int {
		coinNum := 0
		for _, s := range name {
			if s == 'e' || s == 'E' {
				coinNum += 1
			}
			if s == 'i' || s == 'I' {
				coinNum += 2
			}
			if s == 'o' || s == 'O' {
				coinNum += 3
			}
			if s == 'u' || s == 'U' {
				coinNum += 4
			}
		}
		return coinNum
	}
	for _, user := range users {
		// 计算分配的个数
		disNum := calcCoins(user)
		if left <= 0 {
			disNum = 0
		}
		distribution[user] = disNum
		left -= disNum
	}
	return left
}

func main() {
	left := dispatchCoin()
	fmt.Printf("分配如下：%v\n", distribution)
	fmt.Println("剩下：", left)
}
