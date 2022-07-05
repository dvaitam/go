// 141A. Amusing Joke
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s, t, u string
	fmt.Scan(&s, &t, &u)

	m := make(map[rune]int)
	n := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	for _, c := range t {
		m[c]++
	}
	for _, c := range u {
		n[c]++
	}
	if reflect.DeepEqual(m, n) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
