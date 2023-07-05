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
	found := false
	for a := 0; !found && a*1234567 <= 1000000000 && n-a*1234567 >= 0; a++ {
		for b := 0; !found && b*123456 <= 1000000000 && n-a*1234567-b*123456 >= 0; b++ {
			if (n-a*1234567-b*123456)%1234 == 0 {
				found = true
			}
		}
	}
	if found {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
