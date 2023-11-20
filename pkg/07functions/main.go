package main

import "fmt"

// 関数の定義
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// 関数の呼び出し
	greet("Alice")
	greet("Bob")
}
