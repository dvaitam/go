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
		a := make([]int, n+1)
		counts := map[int]int{}
		b := make([]bool, n)
		for i := 1; i <= n; i++ {
			fmt.Fscan(reader, &a[i])
			counts[a[i]]++
			b[a[i]-1] = true
		}
		not_gifted := make([]int, 0)
		for i := 0; i < n; i++ {
			if !b[i] {
				not_gifted = append(not_gifted, i+1)
			}
		}
		//write(f, not_gifted, "\n")
		j := 0
		ans := make([]int, n+1)
		last_user, last_user_gift := -1, -1
		for i := 1; i <= n; i++ {
			if counts[a[i]] > 1 {
				if i == not_gifted[j] {
					if j+1 < len(not_gifted) {
						not_gifted[j], not_gifted[j+1] = not_gifted[j+1], not_gifted[j]
						ans[i] = not_gifted[j]
						last_user, last_user_gift = i, not_gifted[j]
						counts[a[i]]--
						j++
					} else {
						if last_user == -1 {
							ans[i] = a[i]
							//write(f, "something wrong", a, "\n")
						} else {
							ans[i] = last_user_gift
							ans[last_user] = not_gifted[j]
							counts[a[i]]--
							j++
						}

					}
				} else {
					ans[i] = not_gifted[j]
					last_user, last_user_gift = i, not_gifted[j]
					counts[a[i]]--
					j++
				}
			} else {
				ans[i] = a[i]
			}
		}
		write(f, n-len(not_gifted), "\n")
		for i := 1; i <= n; i++ {
			write(f, ans[i], " ")
		}
		write(f, "\n")

	}
}
