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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		if a[i] == 1 && i+1 < n && a[i+1] == 0 && i+2 < n && a[i+2] == 1 {
			ans++
			a[i+2] = 0
		}
	}
	write(f, ans, "\n")
}
