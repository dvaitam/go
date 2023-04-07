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
func get_ans(a int, b int, sq int) int {
	if sq <= 0 {
		return 1000000000
	}
	ans := sq - 1
	ans += a / sq
	if a%sq != 0 {
		ans++
	}
	ans += b / sq
	if b%sq != 0 {
		ans++
	}
	return ans
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
		sq := int(math.Sqrt(float64(a + b)))
		ans := get_ans(a, b, sq)
		for t := sq; true; t++ {
			tmp := get_ans(a, b, t)
			if tmp > ans {
				break
			} else {
				ans = tmp
			}
		}
		for t := sq; true; t-- {
			tmp := get_ans(a, b, t)
			if tmp > ans {
				break
			} else {
				ans = tmp
			}
		}
		write(f, ans, "\n")
	}
}
