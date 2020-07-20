package Practice

import (
	"strings"
)

/**
练习：编写一个回文检测函数，并为其编写单元测试和基准测试，根据测试的结果逐步对其进行优化。
（回文：一个字符串正序和逆序一样，如“Madam,I,mAdam”、“油灯少灯油”等。

需求分析：需要编写一个函数，返回值是bool，参数是字符串，如果参数是回文，返回true
         否则返回false
         参数如果包含英文，英文不区分大小写
*/

/**
 * 回文检测函数
 */
func PalindromeDetec(s string) bool {
	//1.将s转大写后，转换为rune切片
	var chars = []rune(strings.ToUpper(s))
	//2.假设有n个字符，判断第x个字符是否与第(n-x-1)个字符相等
	length := len(chars)
	for i := 0; i < length/2; i++ {
		if chars[i] != chars[length-i-1] {
			return false
		}
	}
	//3.如果全部相等，返回true，有一个不等，返回false
	return true
}
