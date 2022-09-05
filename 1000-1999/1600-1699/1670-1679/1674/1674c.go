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
		var s, t string
		fmt.Fscan(reader, &s, &t)
		if t == "a" {
			fmt.Println(1)
		} else if strings.Contains(t, "a") {
			fmt.Println(-1)
		} else {
			fmt.Println(int64(1) << len(s))
		}
	}
}
