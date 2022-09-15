// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Anything interface {
	int64 | float64 | int | string | float32 | byte
}

func write[K Anything](f *bufio.Writer, a K) {
	f.Write([]byte(fmt.Sprint(a)))
}

func print_two(f *bufio.Writer, n1 int, n2 int) {
	write(f, n1)
	write(f, "\n")
	write(f, n2)
	write(f, "\n")
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var k int
		var s string
		fmt.Fscan(reader, &k, &s)
		m := map[byte]bool{'1': true, '4': true, '6': true, '8': true, '9': true}
		found_two := false
		found_three := false
		found_five := false
		found_seven := false
		ok := false
		for i := 0; i < k; i++ {
			if m[s[i]] {
				ok = true
				print_two(f, 1, int(s[i]-'0'))
				break
			}
		}
		if !ok {

			for i := 0; i < k; i++ {
				if found_two {
					if s[i] == '5' {
						print_two(f, 2, 25)
						break
					}
					if s[i] == '2' {
						print_two(f, 2, 22)
						break
					}
					if s[i] == '7' {
						print_two(f, 2, 27)
						break
					}
				}
				if found_three {
					if s[i] == '2' {
						print_two(f, 2, 32)
						break
					}
					if s[i] == '3' {
						print_two(f, 2, 33)
						break
					}
					if s[i] == '5' {
						print_two(f, 2, 35)
						break
					}

				}
				if found_five {
					if s[i] == '2' {
						print_two(f, 2, 52)
						break
					}
					if s[i] == '5' {
						print_two(f, 2, 55)
						break
					}
					if s[i] == '7' {
						print_two(f, 2, 57)
						break
					}
				}
				if found_seven {
					if s[i] == '2' {
						print_two(f, 2, 72)
						break
					}
					if s[i] == '5' {
						print_two(f, 2, 75)
						break
					}
					if s[i] == '7' {
						print_two(f, 2, 77)
						break
					}
				}
				if s[i] == '2' {
					found_two = true
				}
				if s[i] == '3' {
					found_three = true
				}
				if s[i] == '5' {
					found_five = true
				}
				if s[i] == '7' {
					found_seven = true
				}
			}
		}
	}
}
