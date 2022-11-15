package main

import "fmt"

// 我是单行注释
/*
	我是多行注释
*/

/*
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota是go语言的常量计数器
iota在const关键字出现时将被重置为0(const 内部的第一行之前),
const 中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
iota可以被用作枚举值
*/
const (
	a1 = iota // 默认为0
	a2
	a3
	a4 = "hhh"
)

func main() {
	fmt.Println(a1, a2, a3, a4)

	var a int = 6
	var b int = 100

	// 匿名变量，及多变量声明
	_, c := 1, 2

	// 在Go语言中不需要中间变量即可完成变量交换
	b, a = a, b

	// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	fmt.Println("变量a的内存地址：", &a, "变量b的内存地址：", &b)
	fmt.Println(a, b, c)
}
