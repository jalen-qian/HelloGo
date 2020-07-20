package Split

import (
	"reflect"
	"testing"
	"time"
)

//测试Split

/**
单元测试Split()函数
*/
func TestSplit(t *testing.T) {
	got := Split("我爱你中国", "爱")
	want := []string{"我", "你中国"}
	//对于切片这种引用类型的比较，需要使用reflect.DeepEqual()来比较
	//DeepEqual先比较两个参数类型是否完全相同，再比较内容是否完全相同
	if !reflect.DeepEqual(got, want) {
		t.Errorf("test failed,want:%v got:%v", want, got)
	}
}

/**
当有多个测试用例时，该如何测试呢
不需要写多个函数，我们可以封账测试用例，用map存储起来，并通过t.Run()进行多个子测试
*/
func TestMoreSplit(t *testing.T) {
	//定义表示测试用例的结构体
	type test struct {
		input string   //输入的字符串
		tmp   string   //切割字符串
		want  []string //期望输出
	}
	//定义一个map用来存储多个测试用例
	tests := map[string]test{
		"test1": {"a,b,c,d", ",", []string{"a", "b", "c", "d"}}, //简单
		"test2": {"我爱你爱上你", "爱", []string{"我", "你", "上你"}},      //中文
		"test3": {"我a你就a你", "a你", []string{"我", "就", ""}},       //中英混合
		"test4": {"a你我就a你", "a你", []string{"", "我就", ""}},       //中英混合并且开头结尾都是分割字符串
	}
	//for循环一个一个执行
	for name, test := range tests {
		//t.Run函数，这样写能通过 go test -run=split/test1来单独执行某一个子测试
		t.Run(name, func(t *testing.T) {
			got := Split(test.input, test.tmp)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("test %s failed,want:%v got:%v\n", name, test.want, got)
			}
		})
	}
}

/**
测试覆盖率：
覆盖率就是执行函数时，函数中的语句被执行的数量占函数中总语句数量的比例
覆盖率越高越好，说明函数中没有垃圾代码，不过一般也不需要一定达到100%，达到60%以上就OK了
测试覆盖率使用 go test -cover来进行
*/

/**
基准测试：
基准测试是在一定的负载下，测试函数执行效率（分配内存大小、平均执行时间等）的一种测试
基准测试函数必须是BenchmarkXXX(b *testing.B)这样定义
基准测试时，函数会被执行b.N次，这个N是系统根据当前环境自己调整的，以保障测试的稳定性
*/
func BenchmarkSplit(b *testing.B) {
	//重置时间：
	//假如在测试前需要做一些耗时操作，但是又不希望这部分计入到基准测试的报告中
	time.Sleep(time.Second * 5)
	b.ResetTimer() //调用b.ResetTimer()重新计时
	for i := 0; i < b.N; i++ {
		//不需要测试函数输出是否正确，直接调用
		Split("我爱你，中国爱你", "爱")
	}

}

/**
并行测试：有时候我们需要测试某个函数在并行时的效率，那么可以在基准测试的基础上进行并行测试
基准测试时，函数是在一个goroutine里面跑b.N次，使用并行测试后，系统会将b.N次分给多个
goroutine去跑。
同时也可以自己指定核心数，不指定默认按照GOMAXPROCS来跑

采用并行测试的方法是调用b.RunParallel
*/
func BenchmarkSplitParallel(b *testing.B) {
	//设置并行测试使用的cpu核心个数
	b.SetParallelism(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("我爱你中国，爱你母亲", "爱")
		}
	})
}
