// 1399B. Gifts Fixing
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var T, n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		b := make([]int, n)
		min_a := 1000000000
		min_b := 1000000000
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] < min_a {
				min_a = a[i]
			}
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
			if b[i] < min_b {
				min_b = b[i]
			}
		}
		ans := int64(0)
		for i := 0; i < n; i++ {
			ans += int64(max(a[i]-min_a, b[i]-min_b))
		}
		fmt.Println(ans)

	}
}
