package main

import "fmt"

func DoubleNumber(number *int) {
	*number = *number * 2
}

func main() {
	var number int = 10
	DoubleNumber(&number) // 変数numberを2倍にする処理
	fmt.Println(number)   // 20になる
}
