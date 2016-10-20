package main

import "fmt"

/*
map
map 映射键到值。

map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值。

=======================================================================
map 的文法
map 的文法跟结构体文法相似，不过必须有键名。

=======================================================================
map 的文法（续）
若顶级类型只是一个类型名，你可以在文法的元素中省略它。

=======================================================================
修改 map
在 map m 中插入或修改一个元素：

m[key] = elem
获得元素：

elem = m[key]
删除元素：

delete(m, key)
通过双赋值检测某个键存在：

elem, ok = m[key]
如果 key 在 m 中， ok 为 true。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。

同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。

注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：

elem, ok := m[key]
*/

type Vertex struct {
	Lat, Long float64
}

var m1 map[string]Vertex

func main1() {
	m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])
}

var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main2() {
	fmt.Println(m2)
}

var m3 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main3() {
	fmt.Println(m3)
}

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
