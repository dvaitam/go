// 00
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		ks := make([]int, n)
		ksp := make([]int, 0)
		mp := map[int]bool{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &ks[i])
			if !mp[ks[i]] {
				ksp = append(ksp, ks[i])
				mp[ks[i]] = true
			}
		}
		// if t == 69 {
		// 	fmt.Println(ks)
		// }
		n = len(ksp)
		sort.Ints(ksp)
		for j := 0; j < m; j++ {
			var a, b, c int64
			fmt.Fscan(reader, &a, &b, &c)
			// if t == 69 {
			// 	fmt.Println(a, b, c)
			// }
			if a*c <= 0 {
				write(f, "NO\n")
			} else {
				sq := math.Sqrt(float64(a * c))
				k2 := float64(b) + 2*sq
				k1 := float64(b) - 2*sq
				//write(f, "ks are ", k1, k2, "\n")
				if k1*k2 < 0 {
					k1, k2 = k2, k1
					i := sort.Search(n, func(i int) bool {
						return float64(ksp[i]) >= k1
					})
					//	write(f, "i here", i, "\n")
					if i > 0 && ksp[i-1] >= 0 {
						write(f, "YES\n")

						write(f, ksp[i-1], "\n")
					} else {
						i = sort.Search(n, func(i int) bool {
							return float64(ksp[i]) >= k2
						})
						if i < n && float64(ksp[i]) > k2 && ksp[i] < 0 {
							write(f, "YES\n")
							write(f, ksp[i], "\n")
						} else if i < n-1 && ksp[i+1] < 0 {
							write(f, "YES\n")
							write(f, ksp[i+1], "\n")
						} else {
							write(f, "NO\n")
						}
					}

				} else {
					i := sort.Search(n, func(i int) bool {
						return float64(ksp[i]) >= k1
					})
					//write(f, " i is", i, ksp[i], "\n")
					if i < n {
						if float64(ksp[i]) > k1 && float64(ksp[i]) < k2 {
							write(f, "YES\n")
							write(f, ksp[i], "\n")
						} else if i < n-1 && float64(ksp[i+1]) > k1 && float64(ksp[i+1]) < k2 {
							write(f, "YES\n")
							write(f, ksp[i+1], "\n")
						} else {
							write(f, "NO\n")
						}
					} else {
						write(f, "NO\n")
					}

				}

			}
		}
	}
}
