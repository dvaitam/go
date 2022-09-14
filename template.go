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
	for t := 1; t <= T; t++ {

	}
}
