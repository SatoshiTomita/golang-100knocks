package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文字列の比較
	fmt.Println("比較:", strings.Compare("hello", "world"))

	// 文字列に部分文字列が含まれているか判定
	fmt.Println("含まれているか:", strings.Contains("hello world", "hello"))

	// 文字列を繰り返す
	fmt.Println("繰り返し:", strings.Repeat("hello", 3))

	// 文字列を分割
	fmt.Println("分割:", strings.Split("a b c d e", " "))

	// 文字列を大文字・小文字に変換
	fmt.Println("大文字:", strings.ToUpper("test"))
	fmt.Println("小文字:", strings.ToLower("TEST"))

	// 文字列の前後から指定した文字列を削除
	fmt.Println("前後から削除:", strings.Trim("  test  ", " "))
	fmt.Println("前から削除:", strings.TrimLeft("  test  ", " "))

	// 文字列の前後に指定した文字列を追加
	fmt.Println("前後に追加:", strings.Join([]string{"a", "b", "c"}, ","))
}
