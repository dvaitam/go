// 432A. Choosing Teams
package main

import (
	"fmt"
)

func main() {
	var n, k, curr int
	fmt.Scan(&n, &k)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&curr)
		if 5-curr >= k {
			count++
		}
	}
	fmt.Println(count / 3)
}
