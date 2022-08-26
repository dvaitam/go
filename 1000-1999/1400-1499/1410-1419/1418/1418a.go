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
		var x, y, k int64
		fmt.Fscan(reader, &x, &y, &k)
		dest := (1 + y) * k
		ans := (dest - 1)
		if ans%(x-1) == 0 {
			ans = ans / (x - 1)
		} else {
			ans = ans / (x - 1)
			ans++
		}
		ans += k
		fmt.Println(ans)
	}
}
