package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func sub() {
	fakeStatusChange2()
}
func fakeStatusChange2() {
	// 50文字まで#を表示する
	var hoge = 50
	rand.Seed(time.Now().UnixNano())

	for {

		// 0%～100% まで
		for percent := 0; percent <= 100; percent++ {

			// 最大maxDigits桁に対して、いまの進捗分の#桁数
			progress := percent * hoge / 100

			// ランダムでエラーを発生させる
			seed := rand.Intn(65535)
			hasError := seed < 64 // 65536÷64で1/1024でエラーになる

			// 進捗バーを作る
			marker := "#"
			if hasError {
				marker = "-"
			}
			done := strings.Repeat(marker, progress)
			notyet := strings.Repeat(".", hoge-progress)

			// メッセージの出力
			msg := "\r"
			if hasError {
				msg = "FAIL.\n"
			}
			if percent == 100 {
				msg = "SUCCESS.\n"
			}
			fmt.Printf("%3d%% [%v%v] %v", percent, done, notyet, msg)

			// エラーだったら次のターン
			if hasError {
				break
			}

			// 素敵感の演出
			var osooso = rand.Intn(100)
			time.Sleep(time.Millisecond * time.Duration(osooso))
		}
	}
}
