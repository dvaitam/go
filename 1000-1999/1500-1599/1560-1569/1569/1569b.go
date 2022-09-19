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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = make([]int, n)
		}
		winners := make([]int, 0)
		for i := 0; i < n; i++ {
			if s[i] == '2' {
				winners = append(winners, i)
			}
		}
		if len(winners) == 1 || len(winners) == 2 {
			write(f, "NO\n")
		} else {
			write(f, "YES\n")
			for i := 0; i < len(winners); i++ {
				j, k := winners[i], winners[(i+1)%len(winners)]
				a[j][k] = 1
				a[k][j] = -1
			}
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if i == j {
						write(f, "X")
					} else if a[i][j] == 0 {
						write(f, "=")
					} else if a[i][j] == 1 {
						write(f, "+")
					} else {
						write(f, "-")
					}
				}
				write(f, "\n")
			}
		}

	}
}
