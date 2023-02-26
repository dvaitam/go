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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		k--
		a := make([]int64, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		na := make([]int64, 1)
		na[0] = a[0]
		j := 0
		nk := k
		for i := 1; i < n; i++ {
			if na[j] < 0 && a[i] < 0 || na[j] >= 0 && a[i] >= 0 {
				na[j] += a[i]
			} else {
				na = append(na, a[i])
				j++
			}
			if i == k {
				nk = j
			}
		}

		a = na
		k = nk
		n = len(na)
		i, j := k, k
		possible := true
		health := a[k]
		for true {
			if i <= 0 || j >= n-1 {
				break
			}

			startj := j + 1
			energy_needed := int64(0)
			health_gained := int64(0)
			for startj < n {
				health_gained += a[startj]
				if health_gained < 0 {
					energy_needed -= health_gained
					health_gained = 0
				}
				if health_gained > energy_needed {
					break
				}
				startj++
			}

			if health >= energy_needed && health_gained >= energy_needed {
				j = startj
				health += health_gained - energy_needed
				continue
			}
			starti := i - 1

			energy_needed_left := int64(0)
			health_gained_left := int64(0)
			for starti >= 0 {
				health_gained_left += a[starti]
				if health_gained_left < 0 {
					energy_needed_left -= health_gained_left
					health_gained_left = 0
				}
				if health_gained_left > energy_needed_left {
					break
				}
				starti--
			}

			if health >= energy_needed_left && health_gained_left >= energy_needed_left {
				i = starti
				health += health_gained_left - energy_needed_left
				continue
			}
			//write(f, health_gained, energy_needed, health_gained_left, energy_needed_left, "\n")
			if health < energy_needed_left && health < energy_needed {
				possible = false
			}
			break
		}
		if possible {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
