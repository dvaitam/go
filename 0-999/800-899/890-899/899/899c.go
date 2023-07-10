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
	s := (n * (n + 1)) / 2
	first := make([]int, 0)
	second := make([]int, 0)
	i, j := 1, n
	if s%2 == 1 && n%2 == 1 || n%2 == 1 {
		// if (n-1)%4 == 0 {
		first = append(first, 1)
		// } else {
		// 	second = append(second, 1)

		// }
		i++
	}
	for i+1 < j {
		first = append(first, i)
		first = append(first, j)
		second = append(second, i+1)
		second = append(second, j-1)
		i += 2
		j -= 2
	}
	if i < j {
		// if (n-1)%4 == 0 {
		first = append(first, i)
		second = append(second, j)
		// } else {
		// 	first = append(first, j)
		// 	second = append(second, i)
		// }

	}
	write(f, s%2, "\n")
	write(f, len(first))
	for i := 0; i < len(first); i++ {
		write(f, " ", first[i])
	}
	write(f, "\n")
	// write(f, len(second))
	// for i := 0; i < len(second); i++ {
	// 	write(f, " ", second[i])
	// }
	// write(f, "\n")
}
