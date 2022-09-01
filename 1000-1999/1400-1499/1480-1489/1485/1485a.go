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
		var a, b int
		fmt.Fscan(reader, &a, &b)
		ans := 100
		prob := 0
		if b == 1 {
			b++
			prob++
		}
		for i := 0; i < 32; i++ {
			tmp := prob + i
			tb := b + i
			ta := a
			for ta > 0 {
				ta = ta / tb
				tmp++
			}
			if tmp < ans {
				ans = tmp
			}
		}

		fmt.Println(ans)

	}
}
