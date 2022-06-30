// 110A. Nearly Lucky Number
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	ll := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '4' || s[i] == '7' {
			ll++
		}
	}
	if ll == 4 || ll == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
