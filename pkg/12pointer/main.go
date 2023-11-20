package main

import "fmt"

func main() {
	var num int = 10
	var pointer *int = &num // numのアドレス情報をpointerに格納
	fmt.Println("ポインタの値:", pointer)
	fmt.Println("ポインタ経由の値:", *pointer) // ポインタ経由でnumの値を表示

	*pointer = 20 // ポインタを介してnumの値を更新
	fmt.Println("変更後の値:", num)
}
