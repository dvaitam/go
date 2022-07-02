// 1030A. In Search of an Easy Problem
package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		ans += d
	}
	if ans > 0 {
		fmt.Println("HARD")
	} else {
		fmt.Println("EASY")
	}
}
