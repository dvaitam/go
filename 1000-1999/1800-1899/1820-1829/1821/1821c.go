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
func get_count(a int) int {
	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}
	return 1 + get_count(a/2)
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
		scores := make([]int, 26)
		counts := make([]int, 26)
		for i := 0; i < len(s); i++ {
			for j := 0; j < 26; j++ {
				if j != int(s[i]-'a') {
					counts[j]++
				} else {
					if counts[j] != 0 {
						scores[j] = max(scores[j], get_count(counts[j]))
						counts[j] = 0
					}
				}
			}
		}
		for j := 0; j < 26; j++ {
			if counts[j] != 0 {
				scores[j] = max(scores[j], get_count(counts[j]))
				counts[j] = 0
			}
		}
		ans := scores[0]
		for j := 1; j < 26; j++ {
			ans = min(ans, scores[j])
		}
		write(f, ans, "\n")
	}
}
