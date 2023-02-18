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
func bisect_right(a []int64, x int64) int {
	lo := 0
	hi := len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if x < a[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
func is_prime(n int64, primes []int64) bool {
	for i := 0; i < len(primes); i++ {
		if primes[i]*primes[i] > n {
			return true
		}
		if n%primes[i] == 0 {
			return false
		}
	}
	return true
}
func main() {
	var n int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	primes := make([]int64, 1)
	primes[0] = 2
	for start := int64(3); start*start <= n; start++ {
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
	if n == 2 {
		write(f, "1\n")
	} else if n%2 == 0 {
		write(f, "2\n")
	} else if is_prime(n, primes) {
		write(f, "1\n")
	} else if is_prime(n-2, primes) {
		write(f, "2\n")
	} else {
		write(f, "3\n")
	}
}
