// 996A. Hit the Lottery
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := [5]int{100, 20, 10, 5, 1}
	ans := 0
	for _, v := range a {
		ans += (n / v)
		n = n % v
	}
	fmt.Println(ans)

}
