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
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	//sa := make([]int, n)
	mi := 10000000
	mii := -1
	miii := -1
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		//sa[i] = a[i]
		if a[i] < mi {
			mi = a[i]
			mii = i
		}
	}

	//sort.Ints(sa)
	start := mii
	ok := true
	for i := 0; i < n-1; i++ {
		if a[start%n] <= a[(start+1)%n] {
			start++
		} else {
			ok = false
			break
		}

	}
	if !ok {
		for i := 0; i < n; i++ {
			if a[i] == mi && i != mii && miii == -1 && a[i-1] != mi {
				miii = i
			}
		}
		if miii != -1 {
			start := miii
			ok = true
			for i := 0; i < n-1; i++ {
				if a[start%n] <= a[(start+1)%n] {
					start++
				} else {
					ok = false
					break
				}

			}
			if ok {
				mii = miii
			}
		}
	}

	if ok {
		if mii == 0 {
			write(f, "0\n")
		} else {
			write(f, n-mii, "\n")
		}
	} else {
		write(f, "-1\n")
	}
}
