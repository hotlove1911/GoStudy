// 常量
package main

import "fmt"

// 常量声明
const url = "lifeher.cn"

// 常量省略声明
const (
	c1 = 1
	c2 = 2
	c3 = 3
)

// IOTA
const (
	a = iota
	b
	c
	d
)

// 如果同时声明了多个常量时，如果省略了值则表示和上面的值是一样的。
const (
    name = "name"
    zhangsan
    lisi
    wangwu
)

func main() {
	fmt.Println(url)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(a,b,c,d)
	fmt.Println(zhangsan)
}