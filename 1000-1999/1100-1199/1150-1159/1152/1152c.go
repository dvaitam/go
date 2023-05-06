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
func gcd(a int64, b int64) int64 {
	if b > a {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}
func lcm(a int64, b int64) int64 {
	return a * b / gcd(a, b)
}
func main() {
	var a, b int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &a, &b)
	if a == b {
		write(f, "0\n")
	} else {
		if a < b {
			a, b = b, a
		}
		d := a - b
		if a%d == 0 {
			write(f, "0\n")
		} else {
			if d > b {
				min_lcm := lcm(a, b)
				ans := int64(0)
				for i := int64(1); i*i <= d; i++ {
					if d%i == 0 && d/i >= b {
						if d/i >= b {
							try := d / i
							add := try - b%try
							nxt_lcm := lcm(a+add, b+add)
							if nxt_lcm < min_lcm {
								min_lcm = nxt_lcm
								ans = add
							}
						}
						if i >= b {
							try := i
							add := try - b%try
							nxt_lcm := lcm(a+add, b+add)
							if nxt_lcm < min_lcm {
								min_lcm = nxt_lcm
								ans = add
							}
						}

					}
				}
				write(f, ans, "\n")
			} else {
				write(f, d-b%d, "\n")
			}

		}
	}

}
