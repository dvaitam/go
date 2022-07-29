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
	} else {
		return b
	}
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &p[i])
	}
	ans := a[0] * p[0]
	mi := p[0]
	for i := 1; i < n; i++ {
		mi = min(mi, p[i])
		ans += mi * a[i]
	}
	fmt.Println(ans)
}
