// 546A. Soldier and Bananas
package main

import (
	"fmt"
)

func main() {
	var k, n, w int
	fmt.Scan(&k)
	fmt.Scan(&n)
	fmt.Scan(&w)
	needed := w * (w + 1) * k / 2
	if needed > n {
		fmt.Println(needed - n)
	} else {
		fmt.Println(0)
	}
}
