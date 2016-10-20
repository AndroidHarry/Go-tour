package main

import (
	"fmt"
	"math"
)

/*
练习：循环和函数
作为练习函数和循环的简单途径，用牛顿法实现开方函数。

在这个例子中，牛顿法是通过选择一个初始点 z 然后重复这一过程求 Sqrt(x) 的近似值：


为了做到这个，只需要重复计算 10 次，并且观察不同的值（1，2，3，……）是如何逐步逼近结果的。 然后，修改循环条件，使得当值停止改变（或改变非常小）的时候退出循环。观察迭代次数是否变化。结果与 math.Sqrt 接近吗？

提示：定义并初始化一个浮点值，向其提供一个浮点语法或使用转换：

z := float64(1)
z := 1.0
*/

func Sqrt(x float64) float64 {
	z1 := 1.0
	for z, n := 0.0, 0; n < 10 && math.Abs(z-z1) > 0.001; n++ {
		z = z1
		z1 = z - (z*z-x)/(2*z)
		fmt.Println(z, z1, n, math.Abs(z-z1))
	}
	return z1
}

func main() {
	fmt.Println(Sqrt(2))
}
