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
		ad := make([][]int, n)
		for i := 0; i < n; i++ {
			var k int
			fmt.Fscan(reader, &k)
			ad[i] = make([]int, k)
			for j := 0; j < k; j++ {
				fmt.Fscan(reader, &ad[i][j])
			}
		}
		gr := make([]bool, n)
		br := make([]bool, n)
		for i := 0; i < n; i++ {
			for j := 0; j < len(ad[i]); j++ {
				if !gr[ad[i][j]-1] {
					br[i] = true
					gr[ad[i][j]-1] = true
					break
				}
			}
		}
		ok := true
		gri := -1
		bri := -1
		for i := 0; i < n; i++ {
			if !gr[i] && gri < 0 {
				ok = false
				gri = i + 1
			}
			if !br[i] && bri < 0 {
				ok = false
				bri = i + 1
			}
		}
		if ok {
			write(f, "OPTIMAL\n")
		} else {
			write(f, "IMPROVE\n")
			write(f, bri, gri, "\n")
		}
	}
}
