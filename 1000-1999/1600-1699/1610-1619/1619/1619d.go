// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		var m, n int
		fmt.Fscan(reader, &m, &n)
		p := make([][]int, m)
		uniq := map[int]bool{}
		for i := 0; i < m; i++ {
			p[i] = make([]int, n)

			for j := 0; j < n; j++ {
				fmt.Fscan(reader, &p[i][j])
				uniq[p[i][j]] = true
			}
		}

		keys := make([]int, len(uniq))
		i := 0
		for k := range uniq {
			keys[i] = k
			i++
		}
		sort.Ints(keys)
		start, end := 0, len(keys)

		for start < end {
			mid := (start + end) / 2

			check := keys[mid]
			counts := make([]int, n)
			shops := map[int]int{}
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					if p[i][j] >= check {
						counts[j]++
						shops[i]++
					}
				}
			}
			possible := true
			for i := 0; i < n; i++ {
				if counts[i] == 0 {
					possible = false
					break
				}
			}
			if possible {
				if len(shops) > n-1 {
					number_of_shops := len(shops)
					order := make([]int, len(shops))
					i := 0
					for k := range shops {
						order[i] = k
						i++
					}
					sort.Slice(order, func(i, j int) bool { return shops[order[i]] < shops[order[j]] })
					for i := 0; i < len(order); i++ {
						can_remove := true
						for j := 0; j < n; j++ {
							if p[order[i]][j] >= check && counts[j] == 1 {
								can_remove = false
								break
							}
						}
						if can_remove {
							for j := 0; j < n; j++ {
								if p[order[i]][j] >= check {
									counts[j]--
								}
							}
							number_of_shops--
						}
						if number_of_shops <= n-1 {
							break
						}
					}
					//	write(f, "number of shops", number_of_shops, check, "\n")
					if number_of_shops > n-1 {
						possible = false
					}
				}
			}

			if possible {
				start = mid + 1
			} else {
				end = mid
			}
		}
		write(f, keys[start-1], "\n")
	}
}
