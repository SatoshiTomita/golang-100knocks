package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// フォルダが存在しない場合は作成
	targetDir := "./pkg/26file/out"
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		os.Mkdir(targetDir, 0777)
	}

	// ファイルを作成
	f, err := os.Create(targetDir + "/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// ファイルに書き込み
	_, err = f.WriteString("Hello, world!")
	if err != nil {
		log.Fatal(err)
	}

	// ファイルを読み取り
	content, err := os.ReadFile(targetDir + "/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File content:", string(content))

	// ファイルを削除
	err = os.Remove(targetDir + "/test.txt")
	if err != nil {
		log.Fatal(err)
	}
}
