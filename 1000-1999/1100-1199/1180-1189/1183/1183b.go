// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		a := make([]int, n)
		mi := 1000000000
		mx := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			mi = min(mi, a[i])
			mx = max(mx, a[i])
		}
		ans := mi + k
		if mx-ans > k {
			fmt.Println(-1)
		} else {
			fmt.Println(ans)
		}

	}
}
