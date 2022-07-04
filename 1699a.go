// 1699A. The Third Three Number Problem
package main

import (
	"fmt"
)

func main() {
	var t, n int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Scan(&n)
		if n%2 == 0 {
			fmt.Println(n/2, n/2, 0)
		} else {
			fmt.Println(-1)
		}

	}

}
