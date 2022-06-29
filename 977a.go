// 977A. Wrong Subtraction
package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		if n%10 == 0 {
			n = n / 10
		} else {
			n--
		}
	}
	fmt.Println(n)
}
