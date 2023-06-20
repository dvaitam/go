// 00 730d56143078bb8732213b1d0078a2987903c6a8538e753e3c309b8aba23db53
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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		as, bs, cs := 0, 0, 0
		for i := 0; i < len(s); i++ {
			if s[i] == 'A' {
				as++
			}else if s[i] == 'B' {
				bs++
			}else{
				cs++
			}
		}
		if as+cs == bs {
			write(f, "YES\n")
		}else{
			write(f, "NO\n")
		}
	}
}
