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
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var n int
	var k int
	var s string
	fmt.Fscan(reader, &n, &k, &s)
	b := []byte(s)
	i := 0
	if n == 1 && k > 0 {
		write(f, 0, "\n")
	} else {
		for k > 0 && i < n {
			if i == 0 {
				if b[i] != '1' {
					k--
					b[i] = '1'
				}
			} else {
				if b[i] != '0' {
					k--
					b[i] = '0'
				}
			}
			i++
		}
		write(f, string(b), "\n")
	}

}
