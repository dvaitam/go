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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int64, n)
		pos := make([]int64, 0)
		neg := make([]int64, 0)
		var mi, mx int64
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if i == 0 {
				mi = a[i]
				mx = a[i]
			} else {
				mi = min(mi, a[i])
				mx = max(mx, a[i])
			}
			if a[i] >= 0 {
				pos = append(pos, a[i])
			} else {
				neg = append(neg, a[i])
			}
		}

		// if t == 35050 {
		// 	fmt.Println(a)
		// }
		// if T > 10 {
		// 	continue
		// }
		mxsum := mx - mi
		if mxsum == 0 {
			write(f, "No\n")
		} else {
			write(f, "Yes\n")
			i, j := 0, 0
			total_sum := int64(0)
			for i < len(pos) || j < len(neg) {
				start_sum := int64(0)
				for i < len(pos) && start_sum+pos[i] < mxsum && total_sum+pos[i] < mxsum {
					total_sum += pos[i]
					start_sum += pos[i]
					write(f, pos[i], " ")
					i++
				}
				neg_sum := int64(0)
				for j < len(neg) && -(neg_sum+neg[j]) < mxsum && total_sum+neg[j] >= 0 {
					neg_sum += neg[j]
					total_sum += neg[j]
					write(f, neg[j], " ")
					j++
				}
			}
			write(f, "\n")
		}

	}
}
