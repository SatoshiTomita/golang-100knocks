package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "10"
	num1, _ := strconv.Atoi(str1)
	fmt.Println("数値に変換:", num1+1) // 出力: 11

	str2 := "-10"
	num2, _ := strconv.Atoi(str2)
	fmt.Println("数値に変換:", num2+1) // 出力: -9

	str3 := "hoge"
	num3, err := strconv.Atoi(str3)
	if err != nil {
		fmt.Println("エラー:", err) // 出力: strconv.Atoi: parsing "hoge": invalid syntax
	} else {
		fmt.Println("数値に変換:", num3+1)
	}
}
