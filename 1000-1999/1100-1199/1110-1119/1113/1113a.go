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
	var n, v int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &v)
	ans := 0
	if v >= n-1 {
		ans = n - 1
	} else {
		ans = v
		l := n - 1 - v
		ans += ((l + 1) * (l + 2)) / 2
		ans--
	}
	fmt.Println(ans)

}
