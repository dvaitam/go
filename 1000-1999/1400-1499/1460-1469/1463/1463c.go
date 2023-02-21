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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		time := make([]int64, n+1)
		x := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &time[i], &x[i])
		}
		time[n] = 1000000001 * 3
		curr, nxt := int64(0), int64(0)
		elap_start, elap_end := int64(0), int64(0)
		ans := 0
		positive := true
		for i := 0; i < n; i++ {
			if i == 0 || time[i] >= elap_end {
				curr, nxt = nxt, x[i]
				elap_start, elap_end = time[i], time[i]+abs(nxt-curr)
				if nxt >= curr {
					positive = true
				} else {
					positive = false
				}
			}
			curr_pos := curr + time[i] - elap_start
			end_pos := curr_pos + min(elap_end, time[i+1]) - time[i]
			if !positive {
				curr_pos = curr - (time[i] - elap_start)
				end_pos = curr_pos - (min(elap_end, time[i+1]) - time[i])
			}
			//write(f, "curr, end ", curr_pos, end_pos, "time ", time[i], "\n")
			if x[i] >= min(curr_pos, end_pos) && x[i] <= max(curr_pos, end_pos) {
				//write(f, "winner", curr_pos, end_pos, "\n")
				ans++
			}

		}
		write(f, ans, "\n")
	}
}
