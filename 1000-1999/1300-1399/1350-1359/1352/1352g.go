// 00
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
}

func min[K Number](a K, b K) K {
	if a < b {
		return a
	}
	return b
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
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
		if n <= 3 {
			write(f, "-1\n")
		} else {
			l := list.New()
			l.PushBack(2)
			l.PushBack(4)
			l.PushBack(1)
			l.PushBack(3)
			for k := 5; k <= n; k++ {
				if k%2 == 0 {
					l.PushFront(k)
				} else {
					l.PushBack(k)
				}
			}
			for e := l.Front(); e != nil; e = e.Next() {
				write(f, e.Value.(int))
				write(f, " ")
			}
			write(f, "\n")

		}
	}
}
