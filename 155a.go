// 155A. I_love_%username%
package main

import (
	"fmt"
)

func main() {
	var n, curr int
	fmt.Scan(&n, &curr)
	max := curr
	min := curr
	amazing := 0
	for i := 1; i < n; i++ {
		fmt.Scan(&curr)
		if curr < min {
			min = curr
			amazing++
		}
		if curr > max {
			max = curr
			amazing++
		}
	}
	fmt.Println(amazing)
}
