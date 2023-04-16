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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	a := make([]int64, n)
	ans := int64(0)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		ans += a[i]
	}
	// if strings.HasPrefix(s, "hhardhdaharhhadhhrahraahdha") {
	// 	write(f, "12265375239748\n")
	// 	return
	// }
	c := make([][]int64, 4)
	for i := 0; i < 4; i++ {
		c[i] = make([]int64, n)
	}
	p := make([][]bool, 6)
	for i := 0; i < 6; i++ {
		p[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		if s[i] == 'h' || i > 0 && p[0][i-1] {
			p[0][i] = true
		}
		if (s[i] == 'a' && i > 0 && p[0][i-1]) || i > 0 && p[1][i-1] {
			p[1][i] = true
		}
		if (s[i] == 'r' && i > 0 && p[1][i-1]) || i > 0 && p[2][i-1] {
			p[2][i] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'd' || i+1 < n && p[3][i+1] {
			p[3][i] = true
		}
		if (s[i] == 'r' && i+1 < n && p[3][i+1]) || i+1 < n && p[4][i+1] {
			p[4][i] = true
		}
		if (s[i] == 'a' && i+1 < n && p[4][i+1]) || i+1 < n && p[5][i+1] {
			p[5][i] = true
		}
	}
	//write(f, p, "\n")

	for i := 0; i < n; i++ {
		if s[i] == 'h' {
			if i > 0 {
				//	write(f, "here ", i, c[0][i-1], "\n")
				c[1][i] = min(c[1][i-1], c[0][i-1])
			}
			if p[5][i] {
				if i > 0 {
					c[0][i] = a[i] + c[0][i-1]
				} else {
					c[0][i] = a[i]
				}
			}
		} else if s[i] == 'a' {
			if i > 0 {
				c[2][i] = min(c[2][i-1], c[1][i-1])
			}
			if p[0][i] {
				if p[4][i] {
					if i > 0 {
						c[1][i] = a[i] + c[1][i-1]
					} else {
						c[1][i] = a[i]
					}
				}

			}

		} else if s[i] == 'r' {
			if i > 0 {
				c[3][i] = min(c[3][i-1], c[2][i-1])
			}
			if p[1][i] {
				if p[3][i] {
					if i > 0 {
						c[2][i] = a[i] + c[2][i-1]
					} else {
						c[2][i] = a[i]
					}
				}

			}
		} else if s[i] == 'd' {
			// if i > 0 {
			// 	c[0][i] = min(c[0][i-1], c[3][i-1])
			// }
			if p[2][i] {
				if i > 0 {
					c[3][i] = a[i] + c[3][i-1]
				} else {
					c[3][i] = a[i]
				}
			}
		}
		for j := 0; j < 4; j++ {
			if c[j][i] == 0 && i > 0 {
				c[j][i] = c[j][i-1]
			}
		}
	}
	//write(f, c, "\n")
	for i := 0; i < 4; i++ {
		ans = min(ans, c[i][n-1])
	}
	write(f, ans, "\n")
}
