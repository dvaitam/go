// 1343B. Balanced Array
package main

import (
	"fmt"
)

func main() {
	var T, n int
	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		fmt.Scan(&n)
		if n%4 != 0 {
			fmt.Println("NO")
		} else {
			l := []interface{}{}
			total := 0
			for i := 1; i <= n/2; i++ {
				l = append(l, i*2)
				total += 2 * i
			}
			curr := 1
			for i := 1; i < n/2; i++ {
				l = append(l, curr)
				total -= curr
				curr += 2
			}
			l = append(l, total)
			fmt.Println("YES")
			fmt.Println(l...)
		}
	}

}
