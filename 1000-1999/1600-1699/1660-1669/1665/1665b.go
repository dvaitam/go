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
		var n, curr int
		fmt.Fscan(reader, &n)
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			m[curr]++
		}
		max_count := 0
		for _, v := range m {
			if v > max_count {
				max_count = v
			}
		}
		ans := n - max_count
		for max_count < n {
			ans++
			max_count = max_count * 2
		}
		write_int(f, ans)
		write_string(f, "\n")

	}
}
