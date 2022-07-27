// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		b := make([]int, n)
		m := make(map[int]bool)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			b[i] = a[i]
			m[a[i]] = true
		}
		sort.Sort(sort.Reverse(sort.IntSlice(b)))
		ok := false
		for i := 0; i < n; i++ {
			if a[i] != b[i] {
				ok = true
				break
			}
		}
		if ok || len(m) != n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
