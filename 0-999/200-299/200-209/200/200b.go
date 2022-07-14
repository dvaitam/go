// 200B. Drinks
package main

import (
	"fmt"
)

func main() {
	var n, p int
	fmt.Scan(&n)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&p)
		ans += p
	}

	fmt.Println(float64(ans) / float64(n))

}
