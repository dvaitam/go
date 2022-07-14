// 758A. Holiday Of Equality
package main

import (
	"fmt"
)

func main() {
	var n, curr int
	fmt.Scan(&n)
	total := 0
	max := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&curr)
		total += curr
		if curr > max {
			max = curr
		}
	}
	fmt.Println(max*n - total)
}
