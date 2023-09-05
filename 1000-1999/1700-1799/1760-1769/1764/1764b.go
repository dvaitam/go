// 00
package main

import (
	"bufio"
	"fmt"
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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}
func gcd(a int, b int) int {
	if b > a {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
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
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		sort.Ints(a)
		gd := a[1] - a[0]
		for i := 2; i < n; i++ {
			gd = gcd(gd, a[i]-a[i-1])
		}
		for gd != a[0] {
			a = append(a, gd)
			sort.Ints(a)
			for i := 1; i < n; i++ {
				gd = gcd(gd, a[i]-a[i-1])
			}
		}
		//write(f, a, gd, "\n")
		write(f, 1+(a[len(a)-1]-gd)/gd, "\n")
	}
}
