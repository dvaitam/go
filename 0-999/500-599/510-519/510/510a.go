// 510A. Fox And Snake
package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				fmt.Print("#")
			} else {
				if i%4 == 1 {
					if j == n-1 {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				} else {
					if j == 0 {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
			}
		}
		fmt.Println()
	}

}
