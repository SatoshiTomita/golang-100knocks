package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is less than 5")
	}

	// if文の中で変数を宣言することができる
	if i := 10; i > 5 {
		fmt.Println("i is greater than 5")
	}
}
