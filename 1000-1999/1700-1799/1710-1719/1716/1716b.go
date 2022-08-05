// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		f.Write([]byte(strconv.Itoa(n)))
		f.Write([]byte("\n"))
		for i := 1; i <= n; i++ {
			f.Write([]byte(strconv.Itoa(i)))
			if i != n {
				f.Write([]byte(" "))
			}
		}
		f.Write([]byte("\n"))
		for k := n - 1; k >= 1; k-- {
			for i := 1; i <= n; i++ {
				if i != k {
					f.Write([]byte(strconv.Itoa(i)))
					f.Write([]byte(" "))
				}
			}
			f.Write([]byte(strconv.Itoa(k)))
			f.Write([]byte("\n"))
		}
	}
}
