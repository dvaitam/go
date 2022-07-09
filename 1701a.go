// 1701A. Grass Field
package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		sum := 0
		curr := 0
		for i := 0; i < 4; i++ {
			fmt.Scan(&curr)
			sum += curr
		}
		if sum == 4 {
			fmt.Println(2)
		} else if sum == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
