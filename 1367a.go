// 1367A. Short Substrings
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t int
	var s string
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	t, _ = strconv.Atoi(strings.TrimSpace(tmp))
	for i := 0; i < t; i++ {
		tmp, _ = reader.ReadString('\n')
		s = strings.TrimSpace(tmp)
		n := len(s)/2 + 1
		b := make([]byte, n)
		b[0] = s[0]
		curr := 1
		for j := 1; j < len(s); j += 2 {
			b[curr] = s[j]
			curr++
		}
		fmt.Println(string(b))

	}

}
