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
		var w, h int64
		fmt.Fscan(reader, &w, &h)
		areas := make([]int64, 4)
		for i := 0; i < 4; i++ {
			var k int
			fmt.Fscan(reader, &k)
			xy := make([]int64, k)
			for j := 0; j < k; j++ {
				fmt.Fscan(reader, &xy[j])
			}
			if i == 0 || i == 1 {
				areas[i] = (xy[k-1] - xy[0]) * h
			} else {
				areas[i] = (xy[k-1] - xy[0]) * w
			}
		}
		ans := int64(0)
		for i := 0; i < 4; i++ {
			ans = max(ans, areas[i])
		}
		write(f, ans, "\n")
	}

}
