// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, q int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &q)
	p := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}
	sort.Slice(p, func(i, j int) bool { return p[i] < p[j] })

	for i := 1; i < n; i++ {
		p[i] += p[i-1]
	}
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		ans := p[n-x+y-1]
		if n-x-1 >= 0 {
			ans -= p[n-x-1]
		}
		fmt.Println(ans)
		// add
	}
}
