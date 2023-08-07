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

type Number interface {
	int64 | float64 | int
}

func min[K Number](a K, b K) K {
	if a < b {
		return a
	}
	return b
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	s := make([]string, n)
	t := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}
	a := true
	a90 := true
	a180 := true
	a270 := true
	a277 := true
	ah := true
	av := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] != t[i][j] {
				a = false
			}
			if s[i][j] != t[n-1-j][i] {
				//write(f, i, j, j, n-1-i, "\n")
				a90 = false
			}
			if s[i][j] != t[n-1-i][n-1-j] {
				a180 = false
			}
			if s[i][j] != t[n-1-j][n-1-i] {
				a277 = false
			}
			if s[i][j] != t[j][i] {
				a270 = false
			}
			if s[i][j] != t[i][n-1-j] {
				av = false
			}
			if s[i][j] != t[n-1-i][j] {
				ah = false
			}

		}
	}
	//	write(f, a90, a180, a270, ah, av, "\n")
	ok := a90 || a180 || a270 || ah || av || a || a277
	if ok {
		write(f, "Yes\n")
	} else {
		write(f, "No\n")
	}

}
