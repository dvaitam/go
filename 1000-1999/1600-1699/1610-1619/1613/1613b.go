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
		a := make([]int, n)
		mi := 1000000000
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] < mi {
				mi = a[i]
			}
		}
		total := n / 2
		for i := 0; i < total; i++ {
			if mi != a[i] {
				//fmt.Println(a[i], mi)
				write_int(f, a[i])
				write_string(f, " ")
				write_int(f, mi)
				write_string(f, "\n")
			} else {
				total++
			}
		}

	}
}
