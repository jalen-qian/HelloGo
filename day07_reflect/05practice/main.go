package main

/**
练习：编写代码利用反射实现一个ini文件的解析程序

需求分析：我们这里读取最简单的ini文件，需要识别sections parameters 和 comments
sections 类似 [section1] 单占一行
parameters 类似 aa = bb 每个单独占用一行
comments 类似 ;xxxxx 为注释，需要忽略

1.读取文件
2.识别 sections 和 parameters
3.写入到结构体中

*/
