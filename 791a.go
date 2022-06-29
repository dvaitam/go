// 791A. Bear and Big Brother
package main

import (
	"fmt"
)

func main() {
	var l, b int
	fmt.Scan(&l)
	fmt.Scan(&b)
	ans := 1
	for i := 1; i > 0; i++ {
		l = l * 3
		b = b * 2
		if l > b {
			ans = i
			break
		}
	}
	fmt.Println(ans)
}
