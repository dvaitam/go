// 148A. Insomnia cure
package main

import (
	"fmt"
)

func main() {
	var k, l, m, n, d int
	fmt.Scan(&k, &l, &m, &n, &d)
	ans := 0
	for i := 1; i <= d; i++ {
		if i%k == 0 || i%l == 0 || i%m == 0 || i%n == 0 {
			ans++
		}
	}
	fmt.Println(ans)

}
