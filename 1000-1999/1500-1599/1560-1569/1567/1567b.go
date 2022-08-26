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
	m := make(map[int]int)
	curr := 1
	xor := 0
	for curr <= 300000 {
		m[curr] = xor
		xor = xor ^ curr
		curr++
	}
	for t := 1; t <= T; t++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		if m[a] != b {
			if m[a]^b == a {
				fmt.Println(a + 2)
			} else {
				fmt.Println(a + 1)
			}
		} else {
			fmt.Println(a)
		}
	}
}
