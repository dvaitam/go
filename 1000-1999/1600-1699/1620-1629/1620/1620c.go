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
	uint64 | float64 | int
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
	count uint64
}

func reverse(a []uint64) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var k, x uint64
		var s string
		fmt.Fscan(reader, &n, &k, &x, &s)
		symbols := make([]Symbol, 0)
		for i := 0; i < n; i++ {
			if i == 0 {
				symbols = append(symbols, Symbol{ch: s[i], count: 1})
			} else {
				if s[i] == s[i-1] {
					symbols[len(symbols)-1].count++
				} else {
					symbols = append(symbols, Symbol{ch: s[i], count: 1})
				}
			}
		}
		// if T > 100 {
		// 	if t == 48 {
		// 		write(f, n, k, x, s, "\n")
		// 	} else {
		// 		continue
		// 	}
		// }
		//write(f, symbols, "\n")
		for i := 0; i < len(symbols); i++ {
			if symbols[i].ch == '*' {
				symbols[i].ch = 'b'
				symbols[i].count = symbols[i].count*k + 1
			}
		}
		ans := make([]uint64, len(symbols))
		for i := 0; i < len(symbols); i++ {
			if symbols[i].ch == 'a' {
				ans[i] = symbols[i].count
			}
		}
		//write(f, symbols, "\n")
		vals := make([]uint64, 0)
		powers := make([]uint64, 0)
		curr := uint64(1)
		tx := x
		start := 0
		for i := len(symbols) - 1; i >= 0; i-- {
			if symbols[i].ch == 'b' {
				if tx <= curr*symbols[i].count {
					powers = append(powers, curr)
					start = i
					break
				} else {
					powers = append(powers, curr)
					curr *= symbols[i].count
				}
			}
		}

		j := len(powers) - 1
		for i := start; i < len(symbols); i++ {
			if symbols[i].ch == 'b' {
				vals = append(vals, tx/powers[j])
				tx = tx % powers[j]
				j--
			}
		}
		reverse(vals)
		//write(f, powers, vals, "\n")
		j = 0
		carry := 1
		for i := len(symbols) - 1; i >= 0; i-- {
			if symbols[i].ch == 'b' {
				ans[i] = vals[j]
				if ans[i] < uint64(carry) {
					ans[i] = (ans[i] - uint64(carry) + symbols[i].count) % symbols[i].count
				} else {
					ans[i] = ans[i] - uint64(carry)
					carry = 0
				}
				j++
				if j == len(vals) {
					break
				}
			}
		}
		for i := 0; i < len(symbols); i++ {
			if symbols[i].ch == 'a' {
				for j := uint64(0); j < symbols[i].count; j++ {
					write(f, "a")
				}
			} else {
				for j := uint64(0); j < ans[i]; j++ {
					write(f, "b")
				}
			}
		}
		write(f, "\n")
	}
}
