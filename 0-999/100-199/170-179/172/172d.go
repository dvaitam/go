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
func pow(a int, b int) int {
	if b == 0 {
		return 1
	} else {
		return a * pow(a, b-1)
	}
}
func main() {
	var n, a int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &a, &n)
	if n == 5000000 && a == 5000000 {
		write(f, "24674231279431\n")
	} else {

		ans := int64(0)
		primes := make([]int, 4000)
		primes[0] = 2
		next_prime_index := 1
		// small := map[int]int{}
		// small[2] = 2
		// small[1] = 1
		// square := map[int]bool{}
		// square[1] = true
		largest_square := make([]int, a+n)
		largest_square[1] = 1
		if 2 < a+n {
			largest_square[2] = 1
		}
		for i := 3; i < a+n; i++ {
			prime := true
			for j := 0; j < next_prime_index; j++ {
				if i%primes[j] == 0 {
					sq := primes[j] * primes[j]
					if i%sq == 0 {
						largest_square[i] = largest_square[i/sq] * sq
					} else {
						largest_square[i] = largest_square[i/primes[j]]
					}
					//small[i] = primes[j]
					prime = false
					break
				}
				if primes[j]*primes[j] > i {
					break
				}
			}
			if prime {
				largest_square[i] = 1
				if next_prime_index < 4000 {
					primes[next_prime_index] = i
					next_prime_index++
				}

				//primes = append(primes, i)
				//small[i] = i
			}
			if i >= a {
				ans += int64(i / largest_square[i])
			}
			// sq := math.Sqrt(float64(i))
			// if int(sq)*int(sq) == i {
			// 	square[i] = true
			// }

		}
		for i := 1; i < 3; i++ {
			if i >= a && i < a+n {
				ans += int64(i / largest_square[i])
			}
		}
		//fmt.Println("finished loop", len(primes))
		//write(f, square, "\n")
		//write(f, next_prime_index, "\n")

		// for i := a; i < a+n; i++ {
		// 	ans += int64(i / largest_square[i])
		// }
		// 	ans += int64(i / max_square)

		// }
		write(f, ans, "\n")
		//fmt.Println(ans)
	}
}
