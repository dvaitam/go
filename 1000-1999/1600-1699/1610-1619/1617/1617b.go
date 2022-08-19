// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gcd(a int, b int) int {
	if b > a {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

type Triple struct {
	a, b, c int
}

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
	m := make(map[int]Triple)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		if val, ok := m[n]; ok {
			write_int(f, val.a)
			write_string(f, " ")
			write_int(f, val.b)
			write_string(f, " ")
			write_int(f, val.c)
			write_string(f, "\n")
		} else {
			for i := 2; i > 0; i++ {
				if gcd(n-1-i, i) == 1 {
					write_int(f, i)
					write_string(f, " ")
					write_int(f, n-1-i)
					write_string(f, " ")
					write_int(f, 1)
					write_string(f, "\n")
					m[n] = Triple{a: i, b: n - 1 - i, c: 1}
					break
				}
			}
		}

	}
}
