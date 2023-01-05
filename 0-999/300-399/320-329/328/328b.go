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
	var t, s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &t, &s)
	avail := map[byte]int{}
	req := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '9' {
			avail['6']++
		} else if s[i] == '5' {
			avail['2']++
		} else {
			avail[s[i]]++
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i] == '9' {
			req['6']++
		} else if t[i] == '5' {
			req['2']++
		} else {
			req[t[i]]++
		}
	}
	ans := -1
	for k, v := range req {
		if ans == -1 {
			ans = avail[k] / v
		} else {
			ans = min(avail[k]/v, ans)

		}
	}
	write(f, ans, "\n")
}
