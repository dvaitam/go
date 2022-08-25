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

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		write_int(f, 2)
		write_string(f, "\n")
		a := make([]int, n)
		for i := 1; i <= n; i++ {
			a[i-1] = i
		}
		for len(a) > 1 {
			tmp := len(a)
			write_int(f, a[tmp-1])
			write_string(f, " ")
			write_int(f, a[tmp-2])
			write_string(f, "\n")
			r := (a[tmp-1] + a[tmp-2])
			if r%2 == 0 {
				r = r / 2
			} else {
				r = r/2 + 1
			}
			a[tmp-2] = r
			a = a[:tmp-1]
		}
	}
}
