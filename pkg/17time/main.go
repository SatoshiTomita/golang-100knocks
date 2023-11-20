package main

import (
	"fmt"
	"time"
)

func main() {
	// 現在時刻を取得
	now := time.Now()
	fmt.Println("現在時刻:", now)

	// 年月日を取得
	year := now.Year()
	month := now.Month()
	day := now.Day()
	fmt.Println("年月日:", year, month, day)

	// 明日の日付を取得
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Println("明日の日付:", tomorrow)

	// 特定の時刻を作成
	time1 := time.Date(2023, 5, 12, 15, 30, 0, 0, time.UTC)
	fmt.Println("特定の時刻:", time1)

	// 時刻のフォーマット
	fmt.Println("フォーマット:", now.Format("2006-01-02 15:04:05"))

	// 指定時間だけスリープ
	time.Sleep(2 * time.Second)
	fmt.Println("スリープ終了")

	// 時間の差分を計算
	time2 := time.Date(2023, 5, 12, 15, 30, 0, 0, time.UTC)
	diff := time2.Sub(now)
	fmt.Println("差分:", diff)

	// 今日の0時0分0秒
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Println("今日:", today)
}
