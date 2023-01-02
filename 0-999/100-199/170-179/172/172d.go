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
func pow(a int, b int) int {
	if b == 0 {
		return 1
	} else {
		return a * pow(a, b-1)
	}
}
func main() {
	var n, a int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &a, &n)

	ans := int64(0)
	d := make([]int, 11111111)
	for i := 1; i*i < a+n; i++ {
		ii := i * i
		j := ii
		for j < a+n {
			d[j] = ii
			j += ii
		}
	}
	for i := a; i < a+n; i++ {
		ans += int64(i / d[i])
	}
	write(f, ans, "\n")

}
