// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_factor(n int64) int64 {
	for i := 3; int64(i) <= n/3; i++ {
		if n%int64(i) == 0 {
			return int64(i)
		}
	}
	return n
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, k int64
		fmt.Fscan(reader, &n, &k)
		if n%2 == 0 {
			fmt.Println(n + k*2)
		} else {

			fmt.Println(n + get_factor(n) + (k-1)*2)

		}
	}
}
