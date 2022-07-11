// 1360B. Honest Coach
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var T, n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		fmt.Fscan(reader, &n)
		l := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &l[i])
		}
		sort.Ints(l)
		ans := 10000
		for i := 1; i < n; i++ {
			d := l[i] - l[i-1]
			if d < ans {
				ans = d
			}
		}
		fmt.Println(ans)
	}
}
