package mypackage

import "fmt"

// 大文字で始まる関数は外部から呼び出せる
func SayHello() {
	fmt.Println("Hello from mypackage!")
}
