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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	ok := false
	for i := 4; i < 1000; i++ {
		s := strconv.Itoa(i)
		fit := true
		for j := 0; j < len(s); j++ {
			if s[j] != '4' && s[j] != '7' {
				fit = false
				break
			}
		}
		if !fit {
			continue
		}

		if n%i == 0 {
			ok = true
			break
		}

	}
	if ok {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
