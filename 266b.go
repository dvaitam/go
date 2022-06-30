// 266B. Queue at the School
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, t int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&t)
	fmt.Scan(&s)
	for i := 0; i < t; i++ {
		s = strings.ReplaceAll(s, "BG", "GB")
	}
	fmt.Println(s)

}
