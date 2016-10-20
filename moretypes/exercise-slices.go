package main

import "golang.org/x/tour/pic"

//import "github.com/tour/vendor/golang.org/x/tour/pic"

//import "github.com/Go-zh/tour/vendor/golang.org/x/tour/pic"

/*
实现 Pic 。它返回一个长度为 dy 的 slice，其中每个元素是一个长度为 dx 且元素类型为8位无符号整数的 slice。
当你运行这个程序时， 它会将每个整数作为对应像素的灰度值（好吧，其实是蓝度）并显示这个 slice 所对应的图像。

计算每个像素的灰度值的方法由你决定；几个有意思的选择包括 (x+y)/2、x*y 和 x^y 。

（需要使用循环来分配 [][]uint8 中的每个 []uint8 。）

（使用 uint8(intValue) 来在类型之间进行转换。）
*/

func Pic(dx, dy int) [][]uint8 {
	v := make([][]uint8, dx, dx)
	for x := range v {
		v[x] = make([]uint8, dy, dy)
		for y := range v[x] {
			v[x][y] = uint8((x + y) / 2)
		}
	}
	return v

	//	a := make([][]uint8, dx, dx)
	//	for i := range a {
	//		a[i] = make([]uint8, dy, dy)
	//		for j := range a[i] {
	//			a[i][j] = uint8((i + j) / 2)
	//			// a[i][j] = uint8(i*j)
	//			//a[i][j] = uint8(i ^ j)
	//		}
	//	}

	//	return a
}

func main() {
	pic.Show(Pic)
}
