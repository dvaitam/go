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
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var s string
	fmt.Fscan(reader, &s)
	n := len(s)
	i := 0
	j := n - 1
	count := 0
	for i < j {
		if s[i] != s[j] {
			count++
		}
		i++
		j--
	}
	if count > 1 || (n%2 == 0 && count == 0) {
		write_string(f, "NO\n")
	} else {
		write_string(f, "YES\n")
	}
}
