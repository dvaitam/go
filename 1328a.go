// 1328A. Divisibility Problem
package main

import (
	"fmt"
)

func main() {
	var t, a, b int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Scan(&a, &b)
		if a%b == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(b - (a % b))
		}
	}

}
