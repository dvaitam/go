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
func get_sum(c []int64, i int, j int) int64 {
	if i == 0 {
		return c[j]
	} else {
		return c[j] - c[i-1]
	}
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	c := make([]int64, n)
	first := map[int]int{}
	last := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if i == 0 {
			if a[i] > 0 {
				c[i] = int64(a[i])

			}
		} else {
			if a[i] >= 0 {
				c[i] = c[i-1] + int64(a[i])
			} else {
				c[i] = c[i-1]
			}
		}
		if first[a[i]] == 0 {
			first[a[i]] = i + 1
		}
		last[a[i]] = i + 1
	}
	ans := int64(0)
	start_i := -1
	end_i := -1
	for k, v := range first {

		if last[k] != v {
			sm := get_sum(c, v-1, last[k]-1)
			if k < 0 {
				sm += int64(2 * k)

			}
			if ans == int64(0) || sm > ans {
				ans = sm
				start_i = v - 1
				end_i = last[k] - 1
			}
		}
	}

	cut := make([]int, 0)
	for i := 0; i < n; i++ {
		if (i < start_i || i > end_i) || (a[i] < 0 && i != start_i && i != end_i) {
			cut = append(cut, i+1)
		}
	}

	write(f, ans, len(cut), "\n")
	for i := 0; i < len(cut); i++ {
		write(f, cut[i])
		if i == len(cut)-1 {
			write(f, "\n")
		} else {
			write(f, " ")
		}
	}
}
