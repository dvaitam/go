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

		minh := make([]int64, n)
		rminh := make([]int64, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])

			if i == 0 {
				minh[i] = -a[i]
			} else {
				minh[i] = minh[i-1] - a[i]

			}
			if minh[i] < 0 {
				minh[i] = 0
			}
		}

		for i := n - 1; i >= 0; i-- {

			if i == n-1 {
				rminh[i] = -a[i]
			} else {
				rminh[i] = rminh[i+1] - a[i]
			}
			if rminh[i] < 0 {
				rminh[i] = 0
			}

		}
		possible := false
		health := a[k]

		i, j := k-1, k+1
		if k == 0 || k == n-1 {
			possible = true
		} else {
			if health >= minh[k-1] || health >= rminh[k+1] {
				possible = true
			}
		}

		//	write(f, minh, rminh, i, j, "\n")
		for !possible && i >= 0 && j <= n-1 {
			for a[i] >= 0 {
				health += a[i]
				if health >= minh[i-1] {
					possible = true
					break
				}
				i--
			}
			if possible {
				break
			}
			for a[j] >= 0 {
				health += a[j]
				if health >= rminh[j+1] {
					possible = true
					break
				}
				j++
			}
			if possible {
				break
			}
			//	write(f, health, minh[i], rminh[j], "\n")
			if i >= 0 && health >= minh[i] || j <= n-1 && health >= rminh[j] {
				possible = true
				break
			}

			left_health := health
			starti := i
			max_left_health := int64(0)
			max_left_health_index := i
			for a[starti]+left_health >= 0 {
				left_health += a[starti]
				if left_health >= max_left_health {
					max_left_health = left_health
					max_left_health_index = starti
				}
				starti--
			}

			right_health := health
			startj := j
			max_right_health := int64(0)
			max_right_health_index := j
			for a[startj]+right_health >= 0 {
				right_health += a[startj]
				if right_health >= max_right_health {
					max_right_health = right_health
					max_right_health_index = startj
				}
				startj++
			}
			//write(f, health, max_left_health, max_right_health, max_left_health_index, max_right_health_index, i, j, "\n")

			if i == max_left_health_index && j == max_right_health_index {
				break
			}
			if max_left_health > max_right_health {
				health = max_left_health
				i = max_left_health_index + 1
			} else {
				health = max_right_health
				j = max_right_health_index + 1
			}
		}
		//write(f, a, minh, rminh, "\n")

		if possible {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
