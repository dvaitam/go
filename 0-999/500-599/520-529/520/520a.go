// 520A. Pangram
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	s = strings.ToLower(s)
	m := make(map[byte]bool)
	for i := 0; i < n; i++ {
		m[s[i]] = true
	}
	if len(m) == 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
