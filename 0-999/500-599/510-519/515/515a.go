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
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
func main() {
	var a, b, s int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &a, &b, &s)
	rem := abs(a) + abs(b)
	if s < rem || (s-rem)%2 == 1 {
		write(f, "No\n")
	} else {
		write(f, "Yes\n")
	}
}
