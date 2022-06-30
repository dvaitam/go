// 734A. Anton and Danik
package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	a := 0
	for i := 0; i < n; i++ {
		if s[i] == 'A' {
			a++
		}
	}
	if a > n-a {
		fmt.Println("Anton")
	} else if n-a > a {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}

}
