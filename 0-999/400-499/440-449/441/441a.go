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

func main() {
	var n, k, v int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &v)
	s := make([][]int, n)

	l := []interface{}{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &k)
		s[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &s[i][j])
			if s[i][j] < v {
				if len(l) == 0 || l[len(l)-1] != i+1 {
					l = append(l, i+1)
				}
			}
		}
	}
	write(f, len(l), "\n")
	write(f, l...)
	write(f, "\n")
}
