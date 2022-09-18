// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var s string
	fmt.Fscan(reader, &s)
	s = s[1:]
	n := len(s)
	if strings.Contains(s, "1") || n%2 == 1 {
		write(f, (n/2)+1, "\n")
	} else {
		write(f, n/2, "\n")
	}
}
