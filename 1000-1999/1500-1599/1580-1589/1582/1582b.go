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
		zeros := 0
		ones := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			if curr == 0 {
				zeros++
			}
			if curr == 1 {
				ones++
			}
		}
		ans := int64(ones)
		for zeros > 0 {
			ans = ans * 2
			zeros--
		}
		fmt.Println(ans)

	}
}
