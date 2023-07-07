// 00
package main

import (
	"bufio"
	"fmt"
	"math"
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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	//first_r := -1
	for t := 1; t <= T; t++ {
		var l, r int64
		fmt.Fscan(reader, &l, &r)
		//
		ls := int64(math.Sqrt(float64(l)))
		rs := int64(math.Sqrt(float64(r)))
		if rs*rs > r {
			rs--
		}
		if ls*ls > l {
			ls--
		}
		//write(f, ls*ls, rs*rs, "\n")
		ans := int64(0)
		if ls == rs {
			if rs*rs >= l {
				//	write(f, "again\n")
				ans++
			}
			if rs*rs+rs >= l && rs*rs+rs <= r {
				ans++
			}
			if rs*rs+2*rs >= l && rs*rs+2*rs <= r {
				ans++
			}
		} else {
			if ls*ls < l {
				if ls*ls+ls >= l && ls*ls+ls <= r {
					ans++
				}
				if ls*ls+2*ls >= l && ls*ls+2*ls <= r {
					ans++
				}
				ls++
			}
			if rs*rs >= l {
				//write(f, "again\n")
				ans++
			}
			if rs*rs+rs >= l && rs*rs+rs <= r {
				ans++
			}
			if rs*rs+2*rs >= l && rs*rs+2*rs <= r {
				ans++
			}
		}

		if ls < rs {
			//if rs*rs >= l {
			ans += 3 * (rs - ls)
			// } else {
			// 	if ls < rs-1 {
			// 		ans += 3 * (rs - ls - 1)
			// 	}
			// }
		}
		write(f, ans, "\n")

	}
}
