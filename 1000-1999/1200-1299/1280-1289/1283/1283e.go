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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	mi := map[int]int{}
	mx := map[int]int{}
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)
		mi[x]++
		mx[x]++
	}
	mx_ans := len(mx)
	mi_ans := len(mi)
	//write(f, mx, "\n")
	last_zero := 0
	excess := 0
	for i := 1; i <= n+1; i++ {
		if mx[i] == 0 {
			if excess > 0 {
				if mx[last_zero] == 0 {
					//write(f, "added", last_zero, "\n")
					mx[last_zero]++
					mx_ans++
					excess--
				}
			}
			if excess > 0 {
				//write(f, "added ", i, "\n")
				mx[i]++
				mx_ans++
				excess--
			}
			excess = 0
			last_zero = i
		} else {
			excess += mx[i] - 1
		}
	}
	//write(f, mx, "\n")
	for i := 1; i <= n; i++ {
		if mi[i] > 0 && mi[i+1] > 0 && mi[i+2] > 0 {
			mi[i+1] += mi[i]
			mi[i+1] += mi[i+2]
			mi[i] = 0
			mi[i+2] = 0
			mi_ans -= 2
			i += 2
		} else if mi[i] > 0 && mi[i+2] > 0 {
			mi[i+1] += mi[i]
			mi[i+1] += mi[i+2]
			mi[i] = 0
			mi[i+2] = 0
			mi_ans--
			i += 2
		} else if mi[i] > 0 && mi[i+1] > 0 {
			mi[i] += mi[i+1]
			mi[i+1] = 0
			i++
			mi_ans--
		}
	}
	write(f, mi_ans, mx_ans, "\n")

}
