// 116A. Tram
package main

import (
	"fmt"
)

func main() {
	var n int
	var entry, exit int
	fmt.Scan(&n)
	ans := 0
	in_train := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&exit)
		fmt.Scan(&entry)
		in_train += entry - exit
		if in_train > ans {
			ans = in_train
		}

	}
	fmt.Println(ans)
}
