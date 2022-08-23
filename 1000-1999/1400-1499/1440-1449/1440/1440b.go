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
		fmt.Fscan(reader, &n, &k)
		ans := int64(0)
		a := make([]int64, n*k)
		index := (n / 2) + 1

		for i := 0; i < n*k; i++ {
			fmt.Fscan(reader, &a[i])
		}
		curr := n*k - index
		count := 0
		for count < k {
			ans += a[curr]
			curr = curr - index
			count++
		}
		fmt.Println(ans)
	}
}
