// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}
func gcd(a int, b int) int {
	if b > a {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}
func main() {
	var m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &m)
	cc := make([][]int, 200)
	for i := 0; i < 200; i++ {
		cc[i] = make([]int, 100)
	}
	nu := make([]int, m)
	d := make([]int, m)
	for i := 0; i < m; i++ {
		var s string
		fmt.Fscan(reader, &s)
		pa := strings.Split(s, "/")
		pb := strings.Split(pa[0][1:len(pa[0])-1], "+")
		a, _ := strconv.Atoi(pb[0])
		b, _ := strconv.Atoi(pb[1])
		c, _ := strconv.Atoi(pa[1])
		gd := gcd(a+b, c)
		cc[(a+b)/gd][c/gd]++
		nu[i] = (a + b) / gd
		d[i] = c / gd
	}
	for i := 0; i < m; i++ {
		write(f, cc[nu[i]][d[i]], " ")
	}
	write(f, "\n")
}
