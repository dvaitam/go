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
func lucky(n int64) int64 {
	mx := int64(0)
	mi := int64(9)
	for n > 0 {
		mx = max(n%10, mx)
		mi = min(n%10, mi)
		n = n / 10
	}
	return mx - mi
}
func get_min_max(n int64) (int64, int64) {
	mx := int64(0)
	mi := int64(9)
	for n > 0 {
		mx = max(n%10, mx)
		mi = min(n%10, mi)
		n = n / 10
	}
	return mi, mx
}
func get_single(n int64, d int64) int64 {
	ans := int64(0)
	for n > 1 {
		ans += d
		n = n / 10
		if n > 1 {
			ans = ans * 10
		}
	}
	return ans
}
func digits(n int64) int64 {
	ans := int64(0)
	for n > 0 {
		ans++
		n = n / 10
	}
	return ans
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	first := int64(-1)
	for t := 1; t <= T; t++ {
		var l, r int64
		fmt.Fscan(reader, &l, &r)
		if first == -1 {
			first = l
		}
		if first == 33999931601042809 {
			if t == 136 {
				//fmt.Println("YEs ", l, r)
			}
		}
		ans := l
		ll := lucky(l)
		if lucky(r) < ll {
			ans = r
			ll = lucky(r)
		}
		if digits(l) != digits(r) {
			rem := int64(1)
			l1 := l
			for l1 > 0 {
				l1 = l1 / 10
				rem = rem * 10
			}
			//	write(f, rem, "\n")
			write(f, get_single(rem, 9), "\n")

		} else {

			for l < r {
				l1, r1 := l, r
				rem := int64(1)
				for l1%10 == 0 {
					l1 = l1 / 10
					r1 = r1 / 10
					rem = rem * 10
				}
				if l1 == r1 {
					last_r := (r % rem) / (rem / 10)
					lr := int64(0)
					for lr < last_r {
						tmp := l1*rem + get_single(rem, lr)
						luck := lucky(tmp)
						if luck < ll {
							ans = tmp
							ll = luck
						}
						lr++
					}
					if last_r != 0 {
						l += last_r * (rem / 10)
					} else {
						for last_r == 0 {
							rem = rem / 10
							last_r = (r % rem) / (rem / 10)
						}
						l += 1 * (rem / 10)
					}
				} else {
					lr := l1 % 10
					max_dig := int64(10)
					if l1/10 == r1/10 {
						max_dig = r1 % 10
					}
					for lr < max_dig {
						tmp := (l1/10)*(rem*10) + get_single(rem*10, lr)
						//write(f, "getting single ", l1/10, lr, rem*10, tmp, "\n")
						luck := lucky(tmp)
						if luck < ll {
							ans = tmp
							ll = luck
						}
						lr++
					}
					l += (max_dig - (l1 % 10)) * rem
				}
				//write(f, l, r, "\n")

			}
			write(f, ans, "\n")
		}
	}
}
