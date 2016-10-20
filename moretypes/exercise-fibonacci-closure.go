package main

/*
练习：斐波纳契闭包
现在来通过函数做些有趣的事情。

实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。
*/

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	a, b := 0, 0
	return func() int {
		c := a + b
		if c == 0 {
			c = 1
		}
		a, b = b, c
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
