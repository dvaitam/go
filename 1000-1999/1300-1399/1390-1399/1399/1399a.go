// 1399A. Remove Smallest
package main

import (
	"fmt"
	"sort"
)

func main() {
	var T, n int
	fmt.Scan(&T)
	for t := 1; t <= T; t++ {
		fmt.Scan(&n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		sort.Ints(a)
		possible := true
		for i := 1; i < n; i++ {
			if a[i]-a[i-1] > 1 {
				possible = false
				break
			}
		}
		if possible {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
