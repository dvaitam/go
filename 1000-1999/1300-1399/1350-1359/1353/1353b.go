// 1353B. Two Arrays And Swaps
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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		a := make([]int, n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}
		sort.Ints(a)
		sort.Sort(sort.Reverse(sort.IntSlice(b)))
		ans := 0

		for i := 0; i < n; i++ {
			if b[i] > a[i] && i < k {
				ans += b[i]
			} else {
				ans += a[i]
			}
		}
		fmt.Println(ans)
	}
}
