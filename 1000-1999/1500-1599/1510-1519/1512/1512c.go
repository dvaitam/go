// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
}

func min[K Number](a K, b K) K {
	if a < b {
		return a
	}
	return b
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		var sb string
		fmt.Fscan(reader, &sb)

		s := []byte(sb)
		for i := 0; i < len(s); i++ {
			if s[i] == '1' {
				b--
			} else if s[i] == '0' {
				a--
			}
		}
		ok := true
		i := 0
		j := len(sb) - 1
		for i <= j {
			if s[i] != '?' && s[j] != '?' {
				if s[i] != s[j] {
					ok = false
					break
				}
			}
			i++
			j--
		}
		i = 0
		j = len(sb) - 1
		if ok {
			for i <= j {
				if (s[i] == '?' || s[j] == '?') && (s[i] != s[j]) {
					if s[i] == '?' {
						s[i] = s[j]
						if s[j] == '1' {
							b -= 1
						} else {
							a -= 1
						}
					} else {
						s[j] = s[i]
						if s[j] == '1' {
							b -= 1
						} else {
							a -= 1
						}
					}

				}
				i++
				j--
			}
		}
		i = 0
		j = len(sb) - 1
		if ok {
			for i <= j {
				if s[i] == '?' && s[j] == '?' {
					if i == j {
						if a > 0 {
							s[i] = '0'
							a--
						} else if b > 0 {
							s[i] = '1'
							b--
						}
					} else {
						if a > 1 {
							s[i] = '0'
							s[j] = '0'
							a -= 2
						} else if b > 1 {
							s[i] = '1'
							s[j] = '1'
							b -= 2
						}
					}
				}
				i++
				j--
			}
		}
		if a != 0 || b != 0 {
			ok = false
		}
		if ok {
			write(f, string(s), "\n")
		} else {
			write(f, "-1\n")
		}
	}
}
