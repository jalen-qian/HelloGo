package Split

import (
	"strings"
)

/**
分割字符串函数，s根据tmp来分割，最后返回一个字符串类型的切片
*/
func Split(s, tmp string) (ret []string) {
	//优化：基准测试发现每次执行进行了3次内存分配操作，比较耗时，这里先用make分配内存
	ret = make([]string, 0, strings.Count(s, tmp)+1)
	//找到tmp在s中的第一个位置
	idx := strings.Index(s, tmp)
	//如果能找到
	for idx != -1 {
		//切割字符串s，从第一个找到tmp的位置
		ret = append(ret, s[:idx])
		//将s从tmp的下一个字符切到末尾，并重新赋值
		s = s[idx+len(tmp):]
		idx = strings.Index(s, tmp)
	}
	ret = append(ret, s)
	return
}
