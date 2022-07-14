// 467A. George and Accommodation
package main

import (
	"fmt"
)

func main() {
	var n, p, q int
	fmt.Scan(&n)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&p, &q)
		if q > p+1 {
			ans++
		}
	}
	fmt.Println(ans)
}
