// 00
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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		ans := 0
		ok := true
		for i := n - 1; i > 0; i-- {
			if a[i] == 0 {
				ok = false
				break
			}
			if a[i-1] >= a[i] {
				for a[i-1] >= a[i] {
					a[i-1] = a[i-1] / 2
					ans++
				}
			}
		}
		if ok {
			fmt.Println(ans)
		} else {
			fmt.Println(-1)
		}
	}
}
