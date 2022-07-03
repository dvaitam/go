// 469A. I Wanna Be the Guy
package main

import (
	"fmt"
)

func main() {
	var n, p, a int
	fmt.Scan(&n)
	rounds := map[int]bool{}
	fmt.Scan(&p)
	for i := 0; i < p; i++ {
		fmt.Scan(&a)
		rounds[a] = true
	}
	fmt.Scan(&p)
	for i := 0; i < p; i++ {
		fmt.Scan(&a)
		rounds[a] = true
	}
	if len(rounds) == n {
		fmt.Println("I become the guy.")
	} else {
		fmt.Println("Oh, my keyboard!")
	}

}
