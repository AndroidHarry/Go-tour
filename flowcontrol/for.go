package main

/*
for
Go 只有一种循环结构： for 循环。

基本的 for 循环由三部分组成，它们用分号隔开：

初始化语句：在第一次迭代前执行
条件表达式：在每次迭代前求值
后置语句：在每次迭代的结尾执行
初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。

*注意*：和 C、Java、JavaScript 之类的语言不同，Go 的 for 语句后面没有小括号，大括号 { } 则是必须的。
*/

import "fmt"

func main1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
for（续）
循环初始化语句和后置语句都是可选的。

for 是 Go 的 “while”
基于此可以省略分号：C 的 while 在 Go 中叫做 for 。

无限循环
如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
for {
	}
*/
func main() {
	sum := 1
	for sum < 1000 {
		// for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
