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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int64
		fmt.Fscan(reader, &n)
		f1 := float64(n) / 4
		option1 := int64(math.Sqrt(f1))
		if 4*option1*option1 < n {
			option1++
		}
		option1 = 2*option1 - 1

		option2 := int64((math.Sqrt(float64(n)) - 1) / 2)

		if 4*option2*(option2+1) < n-1 {
			option2++
		}
		option2 = 2 * option2
		write(f, min(option1, option2), "\n")
	}
}
