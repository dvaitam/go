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
		var s string
		fmt.Fscan(reader, &s)
		mx := int64(0)
		count := int64(0)
		first_count := int64(0)
		for i := 0; i < len(s); i++ {
			if i == 0 {
				if s[i] == '1' {
					count++
				}
			} else {
				if s[i] == '0' {
					if s[i-1] == '1' {
						mx = max(mx, count)
						if first_count == 0 {
							first_count = count
						}
						count = 0
					}
				} else {
					count++
				}
			}
		}
		if s[0] == '1' {
			mx = max(mx, count+first_count)
		} else {
			mx = max(mx, count)
		}
		//write(f, "max", mx, "\n")
		if mx == int64(len(s)) {
			write(f, mx*mx, "\n")
		} else {
			side := mx / 2
			if mx%2 == 0 {
				write(f, side*(side+1), "\n")
			} else {
				side++
				write(f, side*side, "\n")
			}
		}

	}
}
