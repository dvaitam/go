// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func write_int(f *bufio.Writer, a int) {
	f.Write([]byte(strconv.Itoa(a)))
}
func write_string(f *bufio.Writer, a string) {
	f.Write([]byte(a))
}
func get_ans(s string) int {
	if len(s) == 1 {
		return 0
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += int(s[i] - '0')
	}
	return 1 + get_ans(strconv.Itoa(ans))
}
func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s)
	fmt.Println(get_ans(s))

}
