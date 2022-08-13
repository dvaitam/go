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
		m := make(map[int]int)
		ans := 0
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]] = a[i]
		}
		i := n - 2
		for i >= 0 {
			if a[i] > a[i+1] {
				break
			}
			i--
		}
		for i >= 0 {
			if m[a[i]] != 0 {
				m[a[i]] = 0
				ans++
			}
			i--
		}

		for i := 0; i < n; i++ {
			a[i] = m[a[i]]
		}

		i = n - 2
		for i >= 0 {
			if a[i] > a[i+1] {
				break
			}
			i--
		}
		for i >= 0 {
			if m[a[i]] != 0 {
				m[a[i]] = 0
				ans++
			}
			i--
		}

		fmt.Println(ans)

	}
}
