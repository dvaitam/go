// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	started := false
	ans := 0
	count := 0
	init := 0
	last_count := 0
	if a[0] == 1 {
		started = true
		count = 1
		i := 0
		for i < n && a[i] == 1 {
			init++
			i++
		}
	}
	if a[n-1] == 1 {
		i := n - 1
		for i >= 0 && a[i] == 1 {
			last_count++
			i--
		}
	}

	for i := 1; i < n; i++ {
		if started {
			if a[i] == 1 {
				count++
			} else {
				started = false
				ans = max(ans, count)
			}
		} else {
			if a[i] == 1 {
				count = 1
				started = true
			}
		}
	}
	ans = max(ans, init+last_count)
	fmt.Println(ans)
}
