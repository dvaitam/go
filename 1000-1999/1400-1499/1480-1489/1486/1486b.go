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
		x := make([]int64, n)
		y := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &x[i], &y[i])
		}
		min_x := int64(1000000000)
		min_y := min_x
		max_x := int64(0)
		max_y := int64(0)
		x_dist := int64(-1)
		y_dist := int64(-1)

		for i := 0; i < n; i++ {
			dist := int64(0)
			for j := 0; j < n; j++ {
				dist += abs(x[i] - x[j])
			}
			if x_dist == -1 {
				x_dist = dist
				min_x = x[i]
				max_x = x[i]
			} else {
				if dist < x_dist {
					min_x = x[i]
					max_x = x[i]
					x_dist = dist
				} else if dist == x_dist {
					min_x = min(min_x, x[i])
					max_x = max(max_x, x[i])
				}
			}
		}
		for i := 0; i < n; i++ {
			dist := int64(0)
			for j := 0; j < n; j++ {
				dist += abs(y[i] - y[j])
			}
			if y_dist == -1 {
				y_dist = dist
				min_y = y[i]
				max_y = y[i]
			} else {
				if dist < y_dist {
					min_y = y[i]
					max_y = y[i]
					y_dist = dist
				} else if dist == y_dist {
					min_y = min(min_y, y[i])
					max_y = max(max_y, y[i])
				}
			}
			//	write(f, "y i ", i, dist, y_dist, "\n")
		}
		//write(f, x_dist, y_dist, min_x, max_x, min_y, max_y, "\n")
		write(f, (max_y-min_y+1)*(max_x-min_x+1), "\n")

	}
}
