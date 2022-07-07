// 472A. Design Tutorial: Learn from Math
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println(4, n-4)
	} else {
		fmt.Println(9, n-9)
	}
}
