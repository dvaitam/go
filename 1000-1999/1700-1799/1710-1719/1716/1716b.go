// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
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
		write(f, n, "\n")

		for i := 1; i <= n; i++ {
			f.Write([]byte(strconv.Itoa(i)))
			if i == n {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}
		for k := n - 1; k >= 1; k-- {
			for i := 1; i <= n; i++ {
				if i != k {
					write(f, i, " ")
				}
			}
			write(f, k, "\n")
		}
	}
}
