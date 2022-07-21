// 276A. Lunch Rush
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func main() {
	// var T int
	reader := bufio.NewReader(os.Stdin)
	// fmt.Fscan(reader, &T)
	// for t := 1; t <= T; t++ {
	var n int
	var k, f, tt int64
	fmt.Fscan(reader, &n, &k, &f, &tt)
	ans := f
	if tt > k {
		ans = f - (tt - k)
	}
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &f, &tt)
		if tt > k {
			ans = max(ans, f-(tt-k))
		} else {
			ans = max(ans, f)
		}
	}
	fmt.Println(ans)
	// }
}
