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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	sum := 0
	for i := 0; i < n; i++ {
		var curr int
		fmt.Fscan(reader, &curr)
		sum += curr
	}
	ans := 0
	for i := 1; i <= 5; i++ {
		if (sum+i)%(n+1) != 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
