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
	primes := make([]int, 0)
	primes = append(primes, 2)
	if n > 2 {
		write(f, 2, "\n")
	} else {
		write(f, 1, "\n")
	}
	for i := 2; i <= n+1; i++ {
		if i == 2 {
			write(f, 1, " ")
		} else {
			prime := true
			for j := 0; j < len(primes) && j*j <= i; j++ {
				if i%primes[j] == 0 {
					prime = false
					break
				}
			}
			if prime {
				primes = append(primes, i)
				write(f, 1, " ")
			} else {
				write(f, 2, " ")
			}
		}
	}
	write(f, "\n")
}
