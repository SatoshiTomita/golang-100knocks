package main

import (
	"errors"
	"fmt"
)

// カスタムエラーを作成
var ErrCustom = errors.New("カスタムエラー")

func main() {
	err := foo()
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		if errors.Is(err, ErrCustom) {
			fmt.Println("カスタムエラーが発生しました")
		}
	}
}

func foo() error {
	// エラーを返す
	return ErrCustom
}
