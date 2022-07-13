// 130A. Again Twenty Five!
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int64
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	fmt.Println(25)
}
