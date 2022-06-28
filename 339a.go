// 339A. Helpful Maths
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "+")
	sort.Strings(s)
	return strings.Join(s, "+")
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Println(SortString(text))

}
