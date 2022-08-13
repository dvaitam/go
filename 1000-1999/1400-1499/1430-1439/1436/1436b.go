// 1436B. Prime Square
package main

import (
	"bufio"
	"fmt"
	"os"
)

var primes map[int]bool

func is_prime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}
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

func main() {
	primes = make(map[int]bool)

	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		b := 1
		for b >= 0 {
			if !is_prime(b) && is_prime(n-1+b) {
				break
			}
			b++
		}
		c := 0
		bb := (n - 1) * b
		for c >= 0 {
			if !is_prime(c) && is_prime(bb+c) {
				break
			}
			c++
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == n-1 && j == n-1 {
					fmt.Print(c)
				} else if i == n-1 || j == n-1 {
					fmt.Print(b)
					if j != n-1 {
						fmt.Print(" ")
					}
				} else {
					fmt.Print(1, " ")
				}
			}
			fmt.Println()
		}
	}
}
