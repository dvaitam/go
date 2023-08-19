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
func in_range(i int, n int) bool {
	return i >= 0 && i < n
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		var x int
		fmt.Fscan(reader, &s, &x)
		n := len(s)
		sb := make([]byte, n)
		for i := 0; i < n; i++ {
			sb[i] = '-'
		}
		ok := true
		for i := 0; i < n; i++ {
			left, right := i-x, i+x
			if in_range(left, n) && !in_range(right, n) {
				if sb[left] == '-' {
					sb[left] = s[i]
				} else {
					if sb[left] != s[i] {
						ok = false
						break
					}
				}
			} else if !in_range(left, n) && in_range(right, n) {
				if sb[right] == '-' {
					sb[right] = s[i]
				} else {
					if sb[right] != s[i] {
						ok = false
						break
					}
				}
			} else if !in_range(left, n) && !in_range(right, n) {
				if s[i] != '0' {
					ok = false
					break
				}
			} else {
				if s[i] == '0' {
					if sb[left] == '1' || sb[right] == '1' {
						ok = false
						break
					}
					sb[left] = '0'
					sb[right] = '0'
				} else {
					if sb[left] == '0' && sb[right] == '0' {
						ok = false
						break
					}

					if sb[left] == '-' {
						sb[left] = '1'
					}

				}
			}

		}
		for i := n - 1; i >= 0; i-- {
			left, right := i-x, i+x
			if in_range(left, n) && !in_range(right, n) {
				if sb[left] == '-' {
					sb[left] = s[i]
				} else {
					if sb[left] != s[i] {
						ok = false
						break
					}
				}
			} else if !in_range(left, n) && in_range(right, n) {
				if sb[right] == '-' {
					sb[right] = s[i]
				} else {
					if sb[right] != s[i] {
						ok = false
						break
					}
				}
			} else if !in_range(left, n) && !in_range(right, n) {
				if s[i] != '0' {
					ok = false
					break
				}
			} else {
				if s[i] == '0' {
					if sb[left] == '1' || sb[right] == '1' {
						ok = false
						break
					}
					sb[left] = '0'
					sb[right] = '0'
				} else {
					if sb[left] == '0' && sb[right] == '0' {
						ok = false
						break
					}
					if sb[right] == '-' {
						sb[right] = '1'
					}
				}
			}

		}
		//	write(f, string(sb), "\n")
		for i := 0; i < n; i++ {
			if sb[i] == '-' {
				sb[i] = '1'
			}
		}
		if ok {
			write(f, string(sb), "\n")
		} else {
			write(f, "-1\n")
		}
	}
}
