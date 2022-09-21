// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	p := make([]int, n+1)
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	for i := 1; i <= n; i++ {
		if ans[i] == 0 {
			visited := make(map[int]bool)
			curr := i
			for !visited[curr] {
				visited[curr] = true
				curr = p[curr]
			}
			if ans[curr] == 0 {
				entry := curr
				for p[entry] != curr {
					ans[entry] = entry
					entry = p[entry]
				}
				ans[entry] = entry
			}

			if i != curr {
				entry := i
				for p[entry] != curr {
					ans[entry] = curr
					entry = p[entry]
				}
				ans[entry] = curr
			}

		}
	}
	for i := 1; i <= n; i++ {
		write(f, ans[i])
		if i == n {
			write(f, "\n")
		} else {
			write(f, " ")
		}
	}
}
