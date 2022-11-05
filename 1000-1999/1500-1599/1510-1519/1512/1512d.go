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
		b := make([]int, n+2)
		high := 0
		shigh := 0
		ps := 0

		for i := 0; i < n+2; i++ {
			fmt.Fscan(reader, &b[i])
			if b[i] <= high && b[i] > shigh {
				shigh = b[i]
			} else if b[i] > high {
				shigh = high
				high = b[i]
			}
			ps += b[i]
		}
		// sort.Ints(b)

		ps -= (high + shigh)
		if ps == high || ps == shigh {
			skip_high := true
			ship_shigh := true
			for i := 0; i < n+2; i++ {
				if skip_high {
					if b[i] == high {
						skip_high = false
						continue
					}
				}
				if ship_shigh {
					if b[i] == shigh {
						ship_shigh = false
						continue
					}
				}
				write(f, b[i], " ")
			}
			write(f, "\n")
		} else {
			rem := ps - (high - shigh)
			ok := false
			skip := -1
			for i := 0; i < n+2; i++ {
				if b[i] == rem && b[i] != shigh && b[i] != high {
					ok = true
					skip = i
					break
				}
			}
			if ok {
				skip_high := true
				for i := 0; i < n+2; i++ {
					if i == skip || b[i] == high {
						if i == skip {
							continue
						} else if skip_high {
							skip_high = false
							continue
						}
					}
					write(f, b[i], " ")
				}
				write(f, "\n")
			} else {
				write(f, "-1\n")
			}
		}
	}
}
