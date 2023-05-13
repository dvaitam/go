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

type Symbol struct {
	ch    byte
	count int
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		counts := make([]int, 0)
		count := 0
		symbols := make([]Symbol, 0)
		for i := 0; i < len(s); i++ {
			if i == 0 {
				count++
				symbols = append(symbols, Symbol{s[i], 1})
			} else {
				if s[i] == s[i-1] {
					symbols[len(symbols)-1].count++
					count++
				} else {
					symbols = append(symbols, Symbol{s[i], 1})
					counts = append(counts, count)
					count = 1
				}
			}
		}
		if count > 0 {
			counts = append(counts, count)
		}

		segments := 0
		sm := 0
		var last *Symbol
		for i := 0; i < len(symbols); i++ {
			if sm&1 == 1 {
				if symbols[i].count > 2 {
					if last == nil {
						last = &symbols[i]
						segments++
					} else {
						if symbols[i].ch != last.ch {
							segments++
						}
						last = &symbols[i]
					}
				}
			} else {
				if symbols[i].count > 1 {
					if last == nil {
						last = &symbols[i]
						segments++
					} else {
						if symbols[i].ch != last.ch {
							segments++
						}
						last = &symbols[i]
					}
				}
			}
			sm += symbols[i].count
		}
		if segments == 0 {
			segments++
		}

		ans := 0
		started := -1
		//write(f, counts, new_counts, again_counts, "\n")
		for i := 0; i < len(counts); i++ {
			if counts[i]&1 == 1 {
				if started != -1 {
					ans += i - started
					started = -1
				} else {
					started = i
				}

			} else {
				if started != -1 {
					if counts[i] != 2 {
					}
				}
			}
		}
		// if segments == 0 {
		// 	segments++
		// }
		write(f, ans, segments, "\n")

	}
}
