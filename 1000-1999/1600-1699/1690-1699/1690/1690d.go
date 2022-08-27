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
	for t := 1; t <= T; t++ {
		var n, k int
		var s string
		fmt.Fscan(reader, &n, &k, &s)
		ans := k
		white := 0
		black := 0
		for i := 0; i < k; i++ {
			if s[i] == 'B' {
				black++
			} else {
				white++
			}
		}
		if white < ans {
			ans = white
		}
		for i := k; i < n; i++ {
			if s[i-k] == 'B' {
				black--
			} else {
				white--
			}
			if s[i] == 'B' {
				black++
			} else {
				white++
			}
			if white < ans {
				ans = white
			}
		}

		fmt.Println(ans)

	}
}
