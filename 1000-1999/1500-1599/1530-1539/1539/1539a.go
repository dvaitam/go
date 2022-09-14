// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Anything interface {
	int64 | float64 | int | string | float32
}

func write[K Anything](f *bufio.Writer, a K) {
	f.Write([]byte(fmt.Sprint(a)))
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for k := 1; k <= T; k++ {
		var n, x, t, ans int64
		fmt.Fscan(reader, &n, &x, &t)
		d := t / x
		if n > d {
			ans = (d*(d+1))/2 + (n-d-1)*d
		} else {
			ans = ((n - 1) * n) / 2
		}
		write(f, ans)
		write(f, "\n")
	}
}
