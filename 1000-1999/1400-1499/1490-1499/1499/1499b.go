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
		var s string
		fmt.Fscan(reader, &s)
		i := 1
		for i < len(s) {
			if s[i] == '1' && s[i-1] == '1' {
				break
			}
			i++
		}
		for i < len(s) {
			if s[i] == '0' && s[i-1] == '0' {
				break
			}
			i++
		}
		if i == len(s) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
