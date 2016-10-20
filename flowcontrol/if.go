package main

import (
	"fmt"
	"math"
)

/*
if
Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
*/
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main1() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
if 的便捷语句
跟 for 一样， if 语句可以在条件之前执行一个简单语句。
由这个语句定义的变量的作用域仅在 if 范围之内。
（在最后的 return 语句处使用 v 看看。）
*/
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main2() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}

/*
if 和 else
在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。
（在 main 的 fmt.Println 调用开始前，两次对 pow 的调用均已执行并返回。）
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
