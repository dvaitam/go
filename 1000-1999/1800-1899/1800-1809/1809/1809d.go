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
		n := len(s)
		ans := int64(n) * 1000000000001
		d := make([]int, n)
		rd := make([]int, n)
		for i := 0; i < n; i++ {
			add := 0
			if s[i] == '1' {
				add = 1
			}
			if i == 0 {
				d[i] = add
			} else {
				d[i] = d[i-1] + add
			}
		}
		for i := n - 1; i >= 0; i-- {
			add := 0
			if s[i] == '0' {
				add = 1
			}
			if i == n-1 {
				rd[i] = add
			} else {
				rd[i] = rd[i+1] + add
			}
		}
		//reverse := int64(0)
		found_one := false
		zeros_after := int64(0)
		for i := 0; i < n; i++ {
			if s[i] == '1' {
				if i+1 < n {
					if s[i+1] == '0' {
						tmp := int64(0)
						if i > 0 {
							tmp += int64(d[i-1]) * 1000000000001
						}
						if i+2 < n {
							tmp += int64(rd[i+2]) * 1000000000001
						}
						tmp += 1000000000000
						ans = min(ans, tmp)
						//reverse++
						//write(f, "here ", i, tmp, "\n")
					}

				}
				if zeros_after == 0 {
					ans = min(ans, int64(rd[i])*1000000000001)
				}
				found_one = true
				tmp := int64(0)
				if i > 0 {
					tmp += int64(d[i-1]) * 1000000000001
				}
				tmp += int64(rd[i]) * 1000000000001
				ans = min(ans, tmp)

			} else {
				if found_one {
					zeros_after++
				}
			}

		}
		ans = min(ans, 1000000000001*int64(rd[0]))
		ans = min(ans, 1000000000001*int64(d[n-1]))
		write(f, ans, "\n")
	}
}
