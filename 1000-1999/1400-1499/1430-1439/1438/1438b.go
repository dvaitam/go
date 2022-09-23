// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, curr int
		fmt.Fscan(reader, &n)
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			m[curr]++
		}
		ok := false
		for _, v := range m {
			if v > 1 {
				ok = true
				break
			}
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
