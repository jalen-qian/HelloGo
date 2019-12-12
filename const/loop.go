package main
import("fmt")

func main(){
	sum := 0
	for i:=1; i <= 100; i++ {
		sum += i;
	}
	fmt.Println("1+2+3+...+100的结果是：", sum);

	j := 0;
	for j < 100 {
		j ++;
		fmt.Print(j, " ");
	}
	fmt.Scanf("sss")
}