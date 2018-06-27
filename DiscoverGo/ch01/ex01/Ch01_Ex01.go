package main

import "fmt"

func main() {
	//타잔송 함수실행
	tjSong(12)
}

// 타잔송 함수
func tjSong(n int) {
	팬티가격 := 10
	칼가격 := 0
	for i := 1; i <= n; i++ {
		팬티가격 = i * 10
		칼가격 = 팬티가격 + 10

		fmt.Println("타잔이 ", 팬티가격, "원 짜리 팬티를 입고", 칼가격, "원 짜리 칼을 차고 노래를 한다.")
	}

}
