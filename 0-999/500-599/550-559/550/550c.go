// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s)
	ok := false
	for i := 0; i <= 1000; i += 8 {
		si := strconv.Itoa(i)
		poss := false
		k := 0
		for j := 0; j < len(s); j++ {
			if s[j] == si[k] {
				k++
			}
			if k == len(si) {
				poss = true
				break
			}
		}
		if poss {
			ok = true
			write(f, "YES\n", si, "\n")
			break
		}
	}
	if !ok {
		write(f, "NO\n")
	}
}
