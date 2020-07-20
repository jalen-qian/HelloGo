package Fib

//计算斐波那契数列的函数
//斐波那契数列：F(1)=1 F(2)=1 F(N) = F(N-1) + F(N-2) N > 2

func Fib(n int) int {
	if n < 1 {
		panic("n不能小于1")
	}
	if n <= 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
