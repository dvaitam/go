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
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = make([]int, m)
		}
		d1 := make(map[int]int)
		d2 := make(map[int]int)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fscan(reader, &a[i][j])
				d1[i+j] += a[i][j]
				d2[i-j] += a[i][j]
			}
		}
		ans := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				tmp := d1[i+j] + d2[i-j] - a[i][j]
				if tmp > ans {
					ans = tmp
				}
			}
		}
		fmt.Println(ans)
		//write_int(f, ans)
		//write_string(f, "\n")
	}
}
