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
		curr := 1
		for curr<<1 < n {
			curr = curr << 1
		}
		for i := curr - 1; i >= 0; i-- {
			write_int(f, i)
			write_string(f, " ")
		}
		for i := curr; i < n; i++ {
			write_int(f, i)
			if i == n-1 {
				write_string(f, "\n")
			} else {
				write_string(f, " ")
			}
		}

	}
}
