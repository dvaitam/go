// 00
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

func main() {
	var n, k int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	sq := math.Sqrt(float64(8*n + 8*k + 9))
	ans1 := (2*n + 3 - int64(sq)) / 2
	ans2 := (2*n + 3 + int64(sq)) / 2
	if ans1 >= 0 {
		write(f, ans1, "\n")
	} else {
		write(f, ans2, "\n")
	}
}
