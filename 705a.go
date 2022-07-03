// 705A. Hulk
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	hate := true
	for i := 0; i < n; i++ {
		if hate {
			fmt.Print("I hate ")
		} else {
			fmt.Print("I love ")
		}
		if i != n-1 {
			fmt.Print("that ")
		}
		hate = !hate
	}
	fmt.Println(" it")

}
