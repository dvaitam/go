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
func main() {
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var n, m, k int64
	fmt.Fscan(reader, &n, &m, &k)
	if k%2 == 0 {
		k = k / 2
		if k == 1 {
			write(f, "YES\n")
			write(f, 0, 0, "\n")
			write(f, 0, m, "\n")
			write(f, n, 0, "\n")
		} else {
			if (n*m)%k != 0 {
				write(f, "NO\n")
			} else {
				write(f, "YES\n")
				gd := gcd(n, k)
				n = n / gd
				k = k / gd
				m = m / k
				write(f, 0, 0, "\n")
				write(f, 0, m, "\n")
				write(f, n, 0, "\n")
			}
		}
	} else {
		if (n*m)%k != 0 {
			write(f, "NO\n")
		} else {
			write(f, "YES\n")
			gd := gcd(n, k)
			if gd > 2 {
				n = (n / gd) * 2
				k = k / gd
				m = m / k
			} else {
				gd = gcd(m, k)
				m = (m / gd) * 2
			}

			write(f, 0, 0, "\n")
			write(f, 0, m, "\n")
			write(f, n, 0, "\n")
		}
	}

}
