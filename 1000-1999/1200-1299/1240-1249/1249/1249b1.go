// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func write_int(f *bufio.Writer, a int) {
	f.Write([]byte(strconv.Itoa(a)))
}
func write_string(f *bufio.Writer, a string) {
	f.Write([]byte(a))
}
func join(d []int, i int, j int) {
	if d[i] != d[j] {
		for k := 1; k < len(d); k++ {
			if d[k] == d[i] {
				d[k] = d[j]
			}
		}
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
		p := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &p[i])
		}
		d := make([]int, n+1)
		for i := 1; i <= n; i++ {
			d[i] = i
		}
		for i := 0; i < n; i++ {
			join(d, i+1, p[i])
		}

		counts := make(map[int]int)
		for i := 1; i < n+1; i++ {
			counts[d[i]]++
		}
		for i := 1; i <= n; i++ {
			write_int(f, counts[d[i]])
			if i == n {
				write_string(f, "\n")
			} else {
				write_string(f, " ")
			}
		}

	}
}
