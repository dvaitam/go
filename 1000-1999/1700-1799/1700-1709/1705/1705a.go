// 1705A. Mark the Photographer
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
		var n, x int
		fmt.Fscan(reader, &n, &x)
		h := make([]int, 2*n)
		for i := 0; i < 2*n; i++ {
			fmt.Fscan(reader, &h[i])
		}
		sort.Ints(h)
		possible := true
		for i := 0; i < n; i++ {
			if h[n+i]-h[i] < x {
				possible = false
				break
			}
		}
		if possible {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
