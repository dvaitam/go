// 271A. Beautiful Year
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	n++
	for n > 0 {
		a, b, c, d := n/1000, (n/100)%10, (n/10)%10, n%10
		if a != b && a != c && a != d && b != c && b != d && c != d {
			fmt.Println(n)
			break
		}
		n++
	}

}
