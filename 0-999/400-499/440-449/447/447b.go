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
	var s string
	var k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &k)
	mx := 0
	w := make([]int, 26)
	for i := 0; i < 26; i++ {
		fmt.Fscan(reader, &w[i])
		if w[i] > mx {
			mx = w[i]
		}
	}
	length := len(s)
	ans := 0
	for i := 0; i < length; i++ {
		ans += w[s[i]-'a'] * (i + 1)
	}
	ans += (mx * (length + 1 + length + k) * k) / 2
	write_int(f, ans)
	write_string(f, "\n")
}
