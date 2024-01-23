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
	a := make([]int, n)
	ones, twos, threes, fours := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if a[i] == 1 {
			ones++
		} else if a[i] == 2 {
			twos++
		} else if a[i] == 3 {
			threes++
		} else {
			fours++
		}
	}
	ans := fours + threes
	ones -= min(ones, threes)
	if twos%2 == 0 {
		ans += twos / 2
		twos = 0
	} else {
		ans += twos / 2
		twos = 1
	}
	//write(f, ans, ones, "\n")
	left := 2*twos + ones
	if left%4 == 0 {
		ans += left / 4
	} else {
		ans += 1 + left/4
	}
	//write(f, left, "\n")
	write(f, ans, "\n")

}
