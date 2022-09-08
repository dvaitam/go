// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		p := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &p[i])
		}
		var s string
		fmt.Fscan(reader, &s)
		count := strings.Count(s, "0")
		wrong_ones := make([]int, 0)
		wrong_zeros := make([]int, 0)
		for i := 0; i < n; i++ {
			if s[i] == '0' && p[i] > count {
				wrong_zeros = append(wrong_zeros, i)
			}
			if s[i] == '1' && p[i] <= count {
				wrong_ones = append(wrong_ones, i)

			}
		}
		// fmt.Println(count, wrong_ones, wrong_zeros)
		for i := 0; i < len(wrong_ones); i++ {
			p[wrong_ones[i]], p[wrong_zeros[i]] = p[wrong_zeros[i]], p[wrong_ones[i]]
		}
		for i := 0; i < n; i++ {
			write_int(f, p[i])
			if i == n-1 {
				write_string(f, "\n")
			} else {
				write_string(f, " ")
			}
		}

	}
}
