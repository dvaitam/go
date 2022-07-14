// 266A. Stones on the Table
package main

import (
	"fmt"
)

func main() {
	var n int
	var text string
	fmt.Scan(&n)
	fmt.Scan(&text)
	unique := 1
	for i := 1; i < len(text); i++ {
		if text[i] != text[i-1] {
			unique += 1
		}
	}
	fmt.Println(len(text) - unique)
}
