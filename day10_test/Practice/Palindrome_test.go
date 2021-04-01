package Practice

import "testing"

/**
单元测试
*/
func TestPalindromeDetec(t *testing.T) {
	// 写多个单元测试用例
	type test struct {
		input string
		want  bool
	}
	// 写5个测试用例
	tests := map[string]test{
		"test1": {"helloOlleH", true},
		"test2": {"我爱你i爱我", false},
		"test3": {"上海自来水来自海上", true},
		"test4": {"油灯少灯油1", false},
		"test5": {"我是zggz是我", true},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := PalindromeDetec(test.input)
			if test.want != got {
				t.Errorf("test failed,want:%v but got:%v\n", test.want, got)
			}
		})
	}
}

/**
基准测试
*/
func BenchmarkPalindromeDetec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeDetec("上海自来水来自海上")
	}
}
