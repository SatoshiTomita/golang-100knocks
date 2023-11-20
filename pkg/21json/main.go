package main

import (
	"encoding/json"
	"fmt"
)

// Person構造体の定義
type Person struct {
	Name     string
	Age      int
	Email    string
	Password string
}

func main() {
	// JSONエンコードの例
	bob := Person{
		Name:     "Bob",
		Age:      25,
		Email:    "bob@example.com",
		Password: "secret",
	}
	jsonBytes, _ := json.Marshal(bob)
	fmt.Println(string(jsonBytes))

	// JSONデコードの例
	jsonString := `{"name":"Alice","age":20}` // バッククォートで囲む
	var alice Person
	json.Unmarshal([]byte(jsonString), &alice)
	fmt.Println(alice.Name) // Alice

	// JSONデコードの例（エラーチェック）
	jsonString = `invalid json`
	var charlie Person
	if err := json.Unmarshal([]byte(jsonString), &charlie); err != nil {
		fmt.Println(err) // 出力: json: unknown field "invalid"
	} else {
		fmt.Println(charlie.Name)
	}
}
