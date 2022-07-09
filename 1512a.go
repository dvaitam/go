// 1512A. Spy Detected!
package main

import (
	"fmt"
)

func main() {
	var T, n, curr, first, second int
	fmt.Scan(&T)
	for t := 1; t <= T; t++ {
		fmt.Scan(&n)
		equal := -1
		ans := -1
		for i := 1; i <= n; i++ {
			if i == 1 {
				fmt.Scan(&first)
			} else if i == 2 {
				fmt.Scan(&second)
				if first == second {
					equal = first
				}
			} else {
				fmt.Scan(&curr)
				if equal == -1 {
					if first == curr {
						equal = first
						ans = 2
					} else {
						equal = second
						ans = 1
					}
				} else {
					if curr != equal {
						ans = i
					}
				}
			}
		}
		fmt.Println(ans)
	}

}
