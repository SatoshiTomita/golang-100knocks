package no4

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello, Go!")
}

func main() {
	// 新しいゴルーチンを生成して関数を実行
	go printHello()

	// 1秒待つ
	time.Sleep(1 * time.Second)
}
