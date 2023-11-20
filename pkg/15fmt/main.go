package main

import (
	"fmt"
	"os"
)

func main() {
	v := 42

	// 文字列を直接表示
	fmt.Print("This", "is", "a", 42, "string\n") // Thisisa42string

	// 別々の文字列をフォーマットして表示
	fmt.Println("Another", "string") // Another string

	// フォーマット指定子を使って表示
	fmt.Printf("The value is %d\n", v) // The value is 42

	// フォーマットしてos.Stdoutに書き込み
	s := "formatted"
	fmt.Fprint(os.Stdout, "Yet", "another", "string\n") // Yetanotherstring
	fmt.Fprintln(os.Stdout, "A", "string")              // A string
	fmt.Fprintf(os.Stdout, "A %s string\n", s)          // A formatted string
}
