// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s, t string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &s, &t)
	m := len(s)
	n := len(t)
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 && s[i] == t[j] {
		i--
		j--
	}
	fmt.Println(i + j + 2)
}
