// 136A. Presents
package main

import (
	"fmt"
)

func main() {
	var n, p int
	fmt.Scan(&n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&p)
		m[p] = i + 1
	}

	for i := 0; i < n; i++ {
		fmt.Print(m[i+1], " ")
	}
	fmt.Println()
}
