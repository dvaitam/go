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
	fmt.Fscan(reader, &T)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	for t := 1; t <= T; t++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i == 0 && j == 0 {
					write_int(f, 0)
				} else {
					if i%4 == 1 || i%4 == 2 {
						if j%4 == 1 || j%4 == 2 {
							write_int(f, 0)
						} else {
							write_int(f, 1)
						}
					} else {
						if j%4 == 1 || j%4 == 2 {
							write_int(f, 1)
						} else {
							write_int(f, 0)
						}

					}
				}
				if j == m-1 {
					write_string(f, "\n")
				} else {
					write_string(f, " ")
				}
			}
		}
	}
}
