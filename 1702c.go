// 1702C. Train and Queries
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
		var n, k, curr int
		fmt.Fscan(reader, &n, &k)
		max := make(map[int]int)
		min := make(map[int]int)
		for i := 1; i <= n; i++ {
			fmt.Fscan(reader, &curr)
			if val, ok := max[curr]; ok {
				if i > val {
					max[curr] = i
				}
			} else {
				max[curr] = i
			}
			if val, ok := min[curr]; ok {
				if i < val {
					min[curr] = i
				}
			} else {
				min[curr] = i
			}
		}
		for i := 1; i <= k; i++ {
			var a, b int
			fmt.Fscan(reader, &a, &b)
			if val1, ok1 := max[b]; ok1 {
				if val2, ok2 := min[a]; ok2 {
					if val1 > val2 {
						fmt.Println("YES")
					} else {
						fmt.Println("NO")
					}
				} else {
					fmt.Println("NO")
				}
			} else {
				fmt.Println("NO")
			}
		}
	}
}
