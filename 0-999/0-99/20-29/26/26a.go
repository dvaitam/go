// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

var primes map[int]bool

func is_prime(n int) bool {
	if primes[n] {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	primes[n] = true
	return true
}
func prime_factors(n int, counter map[int]int) {
	if n == 1 {
		return
	}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if is_prime(i) {
				counter[i]++
				prime_factors(n/i, counter)
			}
		}
	}
}

func main() {
	primes = make(map[int]bool)

	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	ans := 0
	for i := 1; i <= n; i++ {
		counter := make(map[int]int)
		prime_factors(i, counter)
		if len(counter) == 2 {
			ans++
		}

	}
	fmt.Println(ans)

}
