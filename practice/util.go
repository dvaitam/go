// 00
package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func write_int(f *bufio.Writer, a int) {
	f.Write([]byte(strconv.Itoa(a)))
}
func write_string(f *bufio.Writer, a string) {
	f.Write([]byte(a))
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

type Anything interface {
	int64 | float64 | int | string | float32
}

func write[K Anything](f *bufio.Writer, a K) {
	f.Write([]byte(fmt.Sprint(a)))
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(min(4, 5))
}
