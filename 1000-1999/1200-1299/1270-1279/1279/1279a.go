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
		var r, g, b int64
		fmt.Fscan(reader, &r, &g, &b)
		if r > g+b+1 || g > r+b+1 || b > r+g+1 {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}
