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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int64, n)
		b := make([]int64, n)
		area := int64(0)
		max_width := int64(0)
		max_height := int64(0)
		widths := map[int64][]int{}
		heights := map[int64][]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i], &b[i])
			area += a[i] * b[i]
			max_height = max(max_height, a[i])
			max_width = max(max_width, b[i])
			widths[b[i]] = append(widths[b[i]], i)
			heights[a[i]] = append(heights[a[i]], i)
		}
		//write(f, widths, heights, "\n")
		//write(f, area, max_height, max_width, "\n")
		width_possible := false
		if area%max_width == 0 {
			rem_width := max_width
			poss_height := area / max_width
			counts := map[int]bool{}
			counted := 0
			//write(f, "rem ", rem_width, poss_height, "\n")

			for true {
				if len(widths[rem_width]) == 0 {
					//	write(f, "width", rem_width, widths[rem_width], "\n")
					break
				}
				for _, i := range widths[rem_width] {
					if counts[i] {
						continue
					}
					poss_height -= a[i]
					counts[i] = true
					counted++
				}
				//	write(f, "rem ", rem_width, poss_height, "\n")
				//write(f, "heights", poss_height, heights[poss_height], "\n")

				if len(heights[poss_height]) == 0 {
					break
				}
				for _, i := range heights[poss_height] {
					if counts[i] {
						continue
					}
					rem_width -= b[i]
					counts[i] = true
					//write(f, "subtracking ", i, b[i], "\n")
					counted++
				}

			}
			//write(f, "counted ", counted, "\n")
			if counted == n {
				width_possible = true
			}
		}
		height_possible := false
		if area%max_height == 0 {
			rem_width := area / max_height
			poss_height := max_height
			counted := 0
			counts := map[int]bool{}
			for true {
				if len(heights[poss_height]) == 0 {
					break
				}
				for _, i := range heights[poss_height] {
					if counts[i] {
						continue
					}
					counts[i] = true
					rem_width -= b[i]
					counted++
				}
				//write(f, "rem ", rem_width, poss_height, "\n")
				if len(widths[rem_width]) == 0 {
					break
				}
				for _, i := range widths[rem_width] {
					if counts[i] {
						continue
					}
					counts[i] = true
					poss_height -= a[i]
					counted++
				}

			}
			if counted == n {
				height_possible = true
			}
		}
		if height_possible && width_possible && max_height != area/max_width {
			write(f, "2\n")
			write(f, max_height, area/max_height, "\n")
			write(f, area/max_width, max_width, "\n")
		} else if width_possible {
			write(f, "1\n")
			write(f, area/max_width, max_width, "\n")
		} else if height_possible {
			write(f, "1\n")
			write(f, max_height, area/max_height, "\n")
		} else {
			write(f, "something wrong\n")
		}

	}
}
