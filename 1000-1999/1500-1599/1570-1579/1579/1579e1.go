// 00
package main

import (
	"bufio"
	"container/list"
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
		var n, curr int
		fmt.Fscan(reader, &n)
		l := list.New()
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			if l.Len() == 0 {
				l.PushBack(curr)
			} else {
				front := l.Front().Value.(int)
				if curr < front {
					l.PushFront(curr)
				} else {
					l.PushBack(curr)
				}
			}
		}
		for e := l.Front(); e != nil; e = e.Next() {
			write_int(f, e.Value.(int))
			write_string(f, " ")
		}
		write_string(f, "\n")

	}
}
