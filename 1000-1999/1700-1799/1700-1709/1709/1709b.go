// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, p, q int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	h := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}
	da := make([]int64, n)
	dd := make([]int64, n)
	for i := 0; i < n-1; i++ {
		if h[i] > h[i+1] {
			da[i+1] = h[i] - h[i+1]
		}
		da[i+1] += da[i]
	}
	for i := n - 1; i > 0; i-- {
		if h[i] > h[i-1] {
			dd[i-1] = h[i] - h[i-1]
		}
		dd[i-1] += dd[i]
	}
	for i := 0; i < m; i++ {
		damage := int64(0)
		fmt.Fscan(reader, &p, &q)
		if p > q {
			damage = dd[q-1] - dd[p-1]
		} else {
			damage = da[q-1] - da[p-1]
		}

		fmt.Println(damage)
	}
}
