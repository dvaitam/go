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
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	a := make([]int64, 26)
	for i := 0; i < 26; i++ {
		fmt.Fscan(reader, &a[i])
	}
	fmt.Fscan(reader, &s)
	ans := int64(0)
	for i := 0; i < 26; i++ {
		m := map[int64]int64{}
		sm := int64(0)
		for j := 0; j < len(s); j++ {
			sm += a[s[j]-'a']
			if s[j] == 'a'+byte(i) {
				ans += m[sm-a[i]]
				m[sm]++
			}
		}
	}
	write(f, ans, "\n")

}
