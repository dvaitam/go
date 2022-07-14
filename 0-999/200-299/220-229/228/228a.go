// 228A. Is your horseshoe on the other hoof?
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	colors := map[int]bool{}
	colors[a] = true
	colors[b] = true
	colors[c] = true
	colors[d] = true
	fmt.Println(4 - len(colors))
}
