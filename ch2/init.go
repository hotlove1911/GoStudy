// 包的初始化
package main

import "fmt"

func main() {
	
	fmt.Println("内容输出")
}

// 初始化内容的优先级最高
func init(){
	fmt.Println("初始化")
}