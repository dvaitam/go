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
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		u := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if i == 0 || i > 0 && a[i] != a[i-1] {
				u = append(u, a[i])
			}
		}

		ok := true

		inf := 0
		for i := 0; i < len(u); i++ {
			if i > 0 && i < len(u)-1 {
				if u[i] < u[i-1] && u[i] < u[i+1] {
					inf++
				}
			}
		}
		if inf > 0 {
			ok = false
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}

	}
}
