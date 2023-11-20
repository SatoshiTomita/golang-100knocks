package main

import "fmt"

func sendData(channel chan string) {
	// チャネルにデータを送信
	channel <- "Hello, Go!"
}

func main() {
	// チャネルの作成
	messageChannel := make(chan string)

	// ゴルーチンでデータを送信
	go sendData(messageChannel)

	// チャネルからデータを受信して表示
	message := <-messageChannel
	fmt.Println(message)
}
