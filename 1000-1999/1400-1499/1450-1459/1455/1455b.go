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
		var x int
		fmt.Fscan(reader, &x)
		sq := math.Sqrt(float64(8*x + 1))
		ans := int((sq - 1) / 2)
		if (ans*(ans+1))/2 == x {
			write(f, ans, "\n")
		} else {
			if ((ans+1)*(ans+2))/2 == x+1 {
				write(f, ans+2, "\n")
			} else {
				write(f, ans+1, "\n")
			}
		}

	}
}
