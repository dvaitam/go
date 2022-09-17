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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)

		if x == 3 && y > x || x == 2 && y > 3 || x == 1 && y > 1 {
			write(f, "NO\n")
		} else {
			write(f, "YES\n")
		}

	}
}
