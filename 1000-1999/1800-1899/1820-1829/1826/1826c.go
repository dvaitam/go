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

var primes []int64

func get_smallest(n int64) int64 {
	for i := 0; i < len(primes); i++ {
		if primes[i]*primes[i] > n {
			return 1
		}
		if n%primes[i] == 0 {
			return primes[i]
		}
	}
	return 1
}
func main() {
	primes = make([]int64, 1)
	primes[0] = 2
	for start := int64(3); start <= 1000; start++ {
		ok := true
		for i := 0; i < len(primes); i++ {
			if start%primes[i] == 0 {
				ok = false
				break
			}
		}
		if ok {
			primes = append(primes, start)
		}
	}
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int64
		fmt.Fscan(reader, &n, &m)
		small := get_smallest(n)
		if n == 1 {
			write(f, "YES\n")
		} else if m >= n {
			write(f, "NO\n")
		} else if small <= m && small != 1 {
			write(f, "NO\n")
		} else {
			write(f, "YES\n")
		}
	}
}
