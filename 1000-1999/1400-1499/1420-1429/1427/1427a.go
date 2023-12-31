// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func write_int(f *bufio.Writer, a int) {
	f.Write([]byte(strconv.Itoa(a)))
}
func write_string(f *bufio.Writer, a string) {
	f.Write([]byte(a))
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
		s := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			s += a[i]
		}
		if s == 0 {
			write_string(f, "NO\n")
		} else {
			sort.Ints(a)
			ok := true
			s := 0
			for i := 0; i < n; i++ {
				s += a[i]
				if s == 0 {
					ok = false
					break
				}
			}
			if !ok {
				for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
					a[i], a[j] = a[j], a[i]
				}
			}
			write_string(f, "YES\n")
			for i := 0; i < n; i++ {
				write_int(f, a[i])
				if i == n-1 {
					write_string(f, "\n")
				} else {
					write_string(f, " ")
				}
			}
		}
	}
}
