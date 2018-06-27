package main

import "fmt"

func main() {
	fmt.Println(fac1(5))
	fmt.Println()
	fmt.Println(fac2(4))
	fmt.Println()
	fmt.Println(fac3(3))

}

// 팩토리얼 함수
func fac1(n int) int {
	if n <= 0 {
		return 1
	}
	fmt.Println(n)
	return n * fac1(n-1)
}

// 팩토리얼 for문이용 구현
// while문과 같은 역활도 for문이 함
func fac2(n int) int {
	result := 1
	for n > 0 {
		fmt.Println(n)
		result *= n
		n--
	}
	return result
}

// 팩토리얼 for문이용 2
func fac3(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result

}
