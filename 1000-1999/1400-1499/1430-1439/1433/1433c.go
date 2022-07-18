// 1433C. Dominant Piranha
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
		var n int
		fmt.Fscan(reader, &n)
		mx := 0
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			mx = max(mx, a[i])
		}
		ok := false
		ans := 0
		for i := 0; i < n; i++ {
			if a[i] == mx {
				if i > 0 && a[i-1] < a[i] {
					ans = i + 1
					ok = true
					break
				}
				if i < n-1 && a[i+1] < a[i] {
					ans = i + 1
					ok = true
					break
				}
			}
		}
		if ok {
			fmt.Println(ans)
		} else {
			fmt.Println(-1)
		}

	}
}
