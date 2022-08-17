// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	total := 0
	for t := 1; t <= T; t++ {
		var n, q int
		fmt.Fscan(reader, &n, &q)
		a := make([]int, n)
		mx := make([]int, n)
		max_i := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] == n {
				max_i = i
			}
			if i == 0 {
				mx[i] = a[i]
			} else {
				mx[i] = max(a[i], mx[i-1])
			}
		}
		maxed := make([]int, 0)
		for i := 0; i < n; i++ {
			if a[i] == mx[i] {
				maxed = append(maxed, i)
			}
		}
		maxed_wins := make(map[int]int)
		for i := 0; i < len(maxed)-1; i++ {
			maxed_wins[maxed[i]] = maxed[i+1] - maxed[i] - 1
		}
		for i := 0; i < q; i++ {
			total++
			var j, k int
			fmt.Fscan(reader, &j, &k)
			if a[j-1] == n {
				if k <= max_i-1 {
					fmt.Println(0)
				} else if max_i-1 > 0 {
					fmt.Println(k - (max_i - 1))
				} else {
					fmt.Println(k)
				}
			} else {
				if a[j-1] == mx[j-1] {
					wins := maxed_wins[j-1]
					if k <= j-2 {
						fmt.Println(0)
					} else {
						if j > 1 {
							wins++
							fmt.Println(min(k-(j-2), wins))
						} else {
							if k <= wins {
								fmt.Println(k)
							} else {
								fmt.Println(wins)
							}
						}

					}
				} else {
					fmt.Println(0)
				}
			}

		}
	}
}
