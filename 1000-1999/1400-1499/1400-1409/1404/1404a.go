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
		var n, k int
		var s string
		fmt.Fscan(reader, &n, &k, &s)
		sb := []byte(s)
		one := 0
		zero := 0
		ok := true
		for i := 0; i < k; i++ {
			start := sb[i]
			for j := i; j < n; j += k {
				if sb[j] != '?' {
					if start == '?' {
						start = sb[j]
					} else if sb[j] != start {
						ok = false
					}
				}
			}
			if start != '?' {
				for j := i; j < n; j += k {
					sb[j] = start
				}
			}
		}
		for i := 0; i < k; i++ {
			if sb[i] == '1' {
				one++
			} else if sb[i] == '0' {
				zero++
			}
		}
		if abs(one-zero) > k-one-zero {
			ok = false
		} else {
			for i := 0; i < k; i++ {
				if sb[i] == '?' {
					if one < zero {
						sb[i] = '1'
						for j := i; j < n; j += k {
							sb[i] = '1'
						}
						one++
					} else {
						sb[i] = '0'
						for j := i; j < n; j += k {
							sb[i] = '0'
						}
						zero++
					}
				}
			}
		}

		//write(f, string(sb), "\n")
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}

	}
}
