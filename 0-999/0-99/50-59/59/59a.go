// 59A. Word
package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Scan(&x)
	caps := 0
	lower := 0
	for i := 0; i < len(x); i++ {
		if x[i] >= 'a' {
			lower++
		} else {
			caps++
		}

	}
	if caps > lower {
		fmt.Println(strings.ToUpper(x))
	} else {
		fmt.Println(strings.ToLower(x))
	}
}
