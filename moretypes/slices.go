package main

import (
	"fmt"
	"strings"
)

/*
slice
一个 slice 会指向一个序列的值，并且包含了长度信息。

[]T 是一个元素类型为 T 的 slice。

len(s) 返回 slice s 的长度。

====================================================
slice 的 slice
slice 可以包含任意的类型，包括另一个 slice。

====================================================
切片就像数组的引用
切片并不存储任何数据， 它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

与它共享底层数组的切片都会观测到这些修改。

====================================================
切片文法
切片文法类似于没有长度的数组文法。

这是一个数组文法：

[3]bool{true, true, false}
下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

[]bool{true, true, false}

====================================================
切片的默认行为
在进行切片时，你可以利用它的默认行为来忽略上下界。

切片下界的默认值为 0 ，上界则是该切片的长度。

对于数组

var a [10]int
来说，以下切片是等价的：

a[0:10]
a[:10]
a[0:]
a[:]

====================================================
切片的长度与容量
切片拥有 长度 和 容量 。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。

你可以通过重新切片来扩展一个切片，给它提供足够的容量。 试着修改示例程序中的切片操作，向外扩展它的容量，看看会发生什么。

====================================================
nil 切片
切片的零值是 nil 。

nil 切片的长度和容量为 0 且没有底层数组。

====================================================
用 make 创建切片
切面可以用内建函数 make 来创建，这也是你创建动态数组的方式。

make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

a := make([]int, 5)  // len(a)=5
要指定它的容量，需向 make 传入第三个参数：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

====================================================
对 slice 切片
slice 可以重新切片，创建一个新的 slice 值指向相同的数组。

表达式

s[lo:hi]
表示从 lo 到 hi-1 的 slice 元素，含前端，不包含后端。因此

s[lo:lo]
是空的，而

s[lo:lo+1]
有一个元素。

====================================================
向切片追加元素
为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。 内建函数的文档对此函数有详细的介绍。

func append(s []T, vs ...T) []T
append 的第一个参数 s 是一个元素类型为 T 的切片， 其余类型为 T 的值将会追加到该切片的末尾。

append 的结果是一个包含原切片所有元素加上新添加元素的切片。

当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。 返回的切片会指向这个新分配的数组。

（要了解关于切片的更多内容，请阅读文章Go 切片：用法和本质。） https://blog.go-zh.org/go-slices-usage-and-internals

====================================================
range
for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。

当使用 for 循环遍历一个 slice 时，每次迭代 range 将返回两个值。 第一个是当前下标（序号），第二个是该下标所对应元素的一个拷贝。

====================================================
range（续）
可以通过赋值给 _ 来忽略序号和值。

如果只需要索引值，去掉 “ , value ” 的部分即可。
*/

func main1() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}

///////////////////////////////////////////////////////
func main2() {
	// Create a tic-tac-toe board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

///////////////////////////////////////////////////////
func main3() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

///////////////////////////////////////////////////////
func main4() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

///////////////////////////////////////////////////////
func main5() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

///////////////////////////////////////////////////////
func main6() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice6(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice6(s)

	// Extend its length.
	s = s[:4]
	printSlice6(s)

	// Drop its first two values.
	s = s[2:]
	printSlice6(s)
}

func printSlice6(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

///////////////////////////////////////////////////////
func main7() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

///////////////////////////////////////////////////////
func main8() {
	a := make([]int, 5)
	printSlice8("a", a)

	b := make([]int, 0, 5)
	printSlice8("b", b)

	c := b[:2]
	printSlice8("c", c)

	d := c[2:5]
	printSlice8("d", d)
}

func printSlice8(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

///////////////////////////////////////////////////////
func main9() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])
}

///////////////////////////////////////////////////////
func main10() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/******************************************************
range
for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。

当使用 for 循环遍历一个 slice 时，每次迭代 range 将返回两个值。 第一个是当前下标（序号），第二个是该下标所对应元素的一个拷贝。
******************************************************/
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main11() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/******************************************************
range（续）
可以通过赋值给 _ 来忽略序号和值。

如果只需要索引值，去掉 “ , value ” 的部分即可。
******************************************************/
func main12() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
