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
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	c := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}
	s := make([]string, n)
	sr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		sr[i] = Reverse(s[i])
	}
	dpi := make([]int64, n)
	dpr := make([]int64, n)
	dpr[0] = c[0]
	ok := true
	for i := 1; i < n; i++ {
		inc, rev := int64(-1), int64(-1)
		if dpi[i-1] >= 0 {
			if s[i] >= s[i-1] {
				inc = dpi[i-1]
			}
			if sr[i] >= s[i-1] {
				rev = dpi[i-1] + c[i]
			}
		}
		if dpr[i-1] >= 0 {
			if s[i] >= sr[i-1] {
				if inc == -1 {
					inc = dpr[i-1]
				} else {
					inc = min(inc, dpr[i-1])
				}
			}
			if sr[i] >= sr[i-1] {
				if rev == -1 {
					rev = dpr[i-1] + c[i]
				} else {
					rev = min(rev, dpr[i-1]+c[i])
				}
			}
		}
		dpi[i] = inc
		dpr[i] = rev
		if inc == -1 && rev == -1 {
			ok = false
			break
		}
	}
	if !ok {
		write(f, "-1\n")
	} else {
		//write(f, dpi, dpr, "\n")
		if dpi[n-1] >= 0 && dpr[n-1] >= 0 {
			write(f, min(dpi[n-1], dpr[n-1]), "\n")
		} else if dpi[n-1] >= 0 {
			write(f, dpi[n-1], "\n")
		} else {
			write(f, dpr[n-1], "\n")

		}
	}
}
