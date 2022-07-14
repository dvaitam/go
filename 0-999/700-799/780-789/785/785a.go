// 785A. Anton and Polyhedrons
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var s string
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(tmp))
	m := map[string]int{"Tetrahedron": 4, "Cube": 6, "Octahedron": 8, "Dodecahedron": 12, "Icosahedron": 20}

	ans := 0
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSpace(s)
		ans += m[s]
	}
	fmt.Println(ans)

}
