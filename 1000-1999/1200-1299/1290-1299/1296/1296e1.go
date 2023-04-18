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

type Char struct {
	c byte
	i int
}

func other(s string) string {
	if s == "1" {
		return "0"
	} else {
		return "1"
	}
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	chars := make([]Char, n)
	for i := 0; i < n; i++ {
		chars[i] = Char{c: s[i], i: i}
	}
	ok := true
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = "-1"
	}
	for i := 0; i < n; i++ {
		if i > 0 && chars[i].c < chars[i-1].c {
			j := i
			for j > 0 && chars[j].c < chars[j-1].c {
				if ans[chars[j-1].i] == "-1" {
					ans[chars[j-1].i] = "1"
				}
				oth := other(ans[chars[j-1].i])
				if ans[chars[j].i] == "-1" {
					ans[chars[j].i] = oth
				} else {
					if oth != ans[chars[j].i] {
						ok = false
					}
				}
				chars[j], chars[j-1] = chars[j-1], chars[j]
				j--
			}
		}
	}
	if ok {
		write(f, "YES\n")
		for i := 0; i < n; i++ {
			if ans[i] == "-1" {
				write(f, "1")
			} else {
				write(f, ans[i])
			}
		}
		write(f, "\n")
	} else {
		write(f, "NO\n")
	}

}
