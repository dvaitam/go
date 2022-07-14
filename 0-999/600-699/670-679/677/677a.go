// 677A. Vanya and Fence
package main

import (
	"fmt"
)

func main() {
	var n, h, a int
	fmt.Scan(&n, &h)
	w := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a > h {
			w += 2
		} else {
			w++
		}
	}
	fmt.Println(w)

}
