package main

import "fmt"

/*
defer
defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
===========================================================
defer 栈
延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

阅读博文了解更多关于 defer 语句的信息。 http://blog.go-zh.org/defer-panic-and-recover
*/

func main1() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
