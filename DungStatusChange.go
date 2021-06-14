package main

import (
	"fmt"
	"time"
)

func main() {
	fakeStatusChange()
}
func fakeStatusChange() {
	// 50文字まで#を表示する
	hoge := 50
	// カウントアップする変数
	count := 0
	var delay time.Duration
	lineout := 0
	for {
		if hoge != count {
			// 上限まで#を足していく
			fmt.Printf("#")
		} else {
			// 50文字になったら「 done!! + 改行」を出力し
			// 何かが終わった感を演出
			fmt.Println(" done!!")
			// カウントアップをリセット
			count = -1
			lineout++
		}
		count++
		// ランダム値にマッチすると待ち時間をさらに長くする。
		// 何か大きな処理をしているように見えること請け合い
		delay = time.Duration(100) * time.Millisecond
		if 5 < count && count <= 10*lineout {
			time.Sleep(delay)
		}
	}
}
