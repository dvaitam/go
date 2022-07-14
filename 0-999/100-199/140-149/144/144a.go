// 144A. Arrival of the General
package main

import (
	"fmt"
)

func main() {
	var n, val int
	fmt.Scan(&n)
	ans := 0
	min_index := 0
	max_index := 0
	min_val := 101
	max_val := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&val)
		if val > max_val {
			max_val = val
			max_index = i
		}
		if val <= min_val {
			min_val = val
			min_index = i
		}
	}
	ans += max_index
	ans += (n - 1 - min_index)
	if min_index < max_index {
		ans--
	}
	fmt.Println(ans)

}
