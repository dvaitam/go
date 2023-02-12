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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s1, s2 string
		fmt.Fscan(reader, &s1, &s2)
		n1 := len(s1)
		n2 := len(s2)
		nxt := make([][]int, n1+1)
		ok := true
		ans := 1
		nxt[n1] = make([]int, 26)
		for i := n1 - 1; i >= 0; i-- {
			nxt[i] = make([]int, 26)
			copy(nxt[i], nxt[i+1])
			nxt[i][s1[i]-'a'] = i + 1
		}
		j := 0
		curr := 0
		for j < n2 {
			if nxt[curr][s2[j]-'a'] > 0 {
				curr = nxt[curr][s2[j]-'a']
				j++
			} else {
				if curr == 0 {
					ok = false
					break
				}
				curr = 0
				ans++
			}
		}
		if ok {
			write(f, ans, "\n")
		} else {
			write(f, "-1\n")
		}
	}
}
