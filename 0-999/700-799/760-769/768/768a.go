// 768A. Oath of the Night's Watch
package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n, curr int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	mx := 0
	mi := 1000000000
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &curr)
		m[curr]++
		mi = min(mi, curr)
		mx = max(mx, curr)
	}
	ans := 0
	for k, v := range m {
		if k != mi && k != mx {
			ans += v
		}
	}
	fmt.Println(ans)
}
