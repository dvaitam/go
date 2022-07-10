// 1702E. Split Into Two Sets
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, a, b int
		fmt.Fscan(reader, &n)
		max := make(map[int]int)
		min := make(map[int]int)
		possible := true
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a, &b)
			if a == b {
				possible = false
			} else if max[a] == 0 && max[b] == 0 {
				max[a]++
				max[b]++
			} else if min[a] == 0 && min[b] == 0 {
				min[a]++
				min[b]++
			} else {
				possible = false
			}
		}
		if possible {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
