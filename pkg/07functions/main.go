package main

import "fmt"

// 関数の定義
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 引数の型が同じ場合は省略可能
func add(a, b int) int {
	return a + b
}

func main() {
	// 関数の呼び出し
	greet("Alice")
	greet("Bob")
	fmt.Println(add(1, 2))
}
