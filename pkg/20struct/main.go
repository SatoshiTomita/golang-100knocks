package main

// Person構造体の定義
type Person struct {
	Name string
	Age  int
}

func main() {
	// 構造体の初期化
	alice := Person{
		Name: "Alice",
		Age:  12,
	}

	// 構造体のフィールドへのアクセス
	println(alice.Name) // Alice
	println(alice.Age)  // 12
}
