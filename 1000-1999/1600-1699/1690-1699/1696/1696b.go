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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		non_zero := false
		zero_after_non_zero := false
		more := false
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if zero_after_non_zero && a[i] != 0 {
				more = true
			}

			if non_zero && a[i] == 0 {
				zero_after_non_zero = true
			}
			if a[i] != 0 {
				non_zero = true
			}
		}
		if more {
			fmt.Println(2)
		} else {
			if non_zero {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		}

	}
}
