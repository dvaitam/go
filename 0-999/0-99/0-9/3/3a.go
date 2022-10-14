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
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var s1, s2 string
	fmt.Fscan(reader, &s1, &s2)
	x1 := int(s1[0] - 'a')
	y1 := int(s1[1] - '1')
	x2 := int(s2[0] - 'a')
	y2 := int(s2[1] - '1')

	steps := make([]string, 0)

	for x1 != x2 || y1 != y2 {
		if x1 != x2 && y1 != y2 {
			if x1 > x2 && y1 > y2 {
				steps = append(steps, "LD")
				x1 -= 1
				y1 -= 1

			} else if x1 > x2 && y1 < y2 {
				steps = append(steps, "LU")
				x1 -= 1
				y1 += 1

			} else if x1 < x2 && y1 < y2 {
				steps = append(steps, "RU")
				x1 += 1
				y1 += 1
			} else {
				steps = append(steps, "RD")
				x1 += 1
				y1 -= 1

			}
		} else if x1 != x2 {
			if x1 > x2 {
				steps = append(steps, "L")
				x1 -= 1
			} else {
				steps = append(steps, "R")
				x1 += 1
			}
		} else if y1 != y2 {
			if y1 > y2 {
				steps = append(steps, "D")
				y1 -= 1
			} else {
				steps = append(steps, "U")
				y1 += 1
			}
		}
	}
	write(f, len(steps), "\n")
	for i := 0; i < len(steps); i++ {
		write(f, steps[i], "\n")
	}
}
