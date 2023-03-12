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

type Pair struct {
	a, b, i int
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	lesser := make([]Pair, 0)
	greaer := make([]Pair, 0)
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		if a < b {
			lesser = append(lesser, Pair{a: a, b: b, i: i})
		} else {
			greaer = append(greaer, Pair{a: a, b: b, i: i})
		}
	}
	if len(lesser) > len(greaer) {
		sort.Slice(lesser, func(i, j int) bool { return lesser[i].b > lesser[j].b })
		write(f, len(lesser), "\n")
		for i := 0; i < len(lesser); i++ {
			write(f, lesser[i].i, " ")
		}
		write(f, "\n")
	} else {
		sort.Slice(greaer, func(i, j int) bool { return greaer[i].b < greaer[j].b })
		write(f, len(greaer), "\n")
		for i := 0; i < len(greaer); i++ {
			write(f, greaer[i].i, " ")
		}
		write(f, "\n")
	}

}
