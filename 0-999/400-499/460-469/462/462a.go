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
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	ok := true
	for i := 0; i < n; i++ {
		if !ok {
			break
		}
		for j := 0; j < n; j++ {
			count := 0
			if j+1 < n && s[i][j+1] == 'o' {
				count++
			}
			if i+1 < n && s[i+1][j] == 'o' {
				count++
			}
			if j-1 >= 0 && s[i][j-1] == 'o' {
				count++
			}
			if i-1 >= 0 && s[i-1][j] == 'o' {
				count++
			}
			if count%2 == 1 {
				ok = false
				break
			}
		}
	}
	if ok {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
