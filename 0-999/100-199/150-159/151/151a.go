// 151A. Soft Drinking
package main

import (
	"fmt"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var n, k, l, c, d, p, nl, np int
	fmt.Scan(&n, &k, &l, &c, &d, &p, &nl, &np)
	count1 := (k * l) / (nl * n)
	count2 := (c * d) / n
	count3 := p / (np * n)
	fmt.Println(min(count1, min(count2, count3)))
}
