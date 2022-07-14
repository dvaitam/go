// 581A. Vasya the Hipster
package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if b > a {
		a, b = b, a
	}
	fmt.Println(b, (a-b)/2)

}
