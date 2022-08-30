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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	ok := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				if s[i][j] != s[0][0] {
					ok = false
				}
			}
		}
	}
	if s[0][0] == s[0][1] {
		ok = false
	}
	if ok {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if !(i == j || i+j == n-1) {
					if s[i][j] != s[0][1] {
						ok = false
					}
				}
			}
		}
	}
	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
