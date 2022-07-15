// 320A - Magic Numbers
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &s)
	magic := true
	if s[0] != '1' {
		magic = false
	}
	for i := 0; i < len(s)-1; i++ {
		if i == 0 && s[i] != '1' {
			magic = false
			break
		}
		if s[i] != '1' && s[i] != '4' {
			magic = false
			break
		}
		if s[i+1] != '1' && s[i+1] != '4' {
			magic = false
			break
		}
		if s[i] == '4' {
			if s[i-1] == '4' && s[i+1] != '1' {
				magic = false
				break
			}
		}
	}
	if magic {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
