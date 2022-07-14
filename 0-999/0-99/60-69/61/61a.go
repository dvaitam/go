// 61A. Ultra-Fast Mathematician
package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	ans := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			ans[i] = '0'
		} else {
			ans[i] = '1'
		}
	}
	fmt.Println(string(ans))

}
