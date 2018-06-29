package main

import (
	"fmt"
)

func main() {
	pibona(10)
}

// 피보나치수열
func pibona(n int) {
	p, q := 0, 1
	fmt.Print(p)
	for i := 1; i <= n; i++ {
		p, q = q, p+q
		fmt.Print(",", p)
	}
}
