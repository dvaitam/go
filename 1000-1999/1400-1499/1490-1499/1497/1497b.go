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
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([]int, n)
		mp := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			mp[a[i]%m]++
		}

		counted := 0
		ans := 0
		for k, v := range mp {
			if k == 0 || k == m-k {
				counted += v
				ans++
			} else if k < m-k {
				if mp[k] > 0 && mp[m-k] > 0 {
					mx := max(mp[k], mp[m-k])
					mi := min(mp[k], mp[m-k])
					if mi != mx {
						counted += mi + mi + 1
						ans++
					} else {
						counted += mi * 2
						ans++
					}
				}
			}
		}
		ans += (n - counted)
		write(f, ans, "\n")

	}
}
