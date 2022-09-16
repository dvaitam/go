// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
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
		s := make([]int, n)
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &s[i])
			m[s[i]]++
		}

		ok := true
		for _, v := range m {
			if v == 1 {
				ok = false
				break
			}
		}
		if ok {
			keys := make([]int, len(m))
			i := 0
			for k := range m {
				keys[i] = k
				i++
			}
			sort.Ints(keys)

			start := 1
			for _, kk := range keys {
				v := m[kk]
				write(f, start+v-1, " ")
				for k := start; k < start+v-1; k++ {
					write(f, k, " ")
				}
				start += v
			}
			// print line
			write(f, "\n")

		} else {
			write(f, -1, "\n")
		}
	}
}
