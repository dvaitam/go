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

var found bool
var cache map[int]bool

func check(u int, v int) {
	if found {
		return
	}
	if u == v {
		found = true
		return
	}
	if u > v {
		return
	}
	if cache[v] {
		return
	}
	cache[v] = true
	// if v%2 == 0 {
	// 	check(u, v/2)
	// 	return
	// }
	for i := 30; i >= 0; i-- {
		curr := 1 << i
		if curr > v {
			continue
		}
		if curr|v == curr+v && !found && v-curr >= u {
			check(u, v-curr)
			break
		}
	}
	//fmt.Println("checked", u, v)

}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		// if T > 600 {
		// 	if t == 641 {
		// 		write(f, u, v, "\n")
		// 	} else {
		// 		continue
		// 	}
		// }
		found = false
		cache = map[int]bool{}
		check(u, v)
		if found {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
