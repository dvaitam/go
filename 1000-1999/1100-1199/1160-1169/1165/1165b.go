// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func write_int(f *bufio.Writer, a int) {
	f.Write([]byte(strconv.Itoa(a)))
}
func write_string(f *bufio.Writer, a string) {
	f.Write([]byte(a))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	ans := 0
	curr := 1
	for i := 0; i < n; i++ {
		if curr <= a[i] {
			ans++
			curr++
		}
	}
	write_int(f, ans)
	write_string(f, "\n")

}
