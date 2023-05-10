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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}

type Symbol struct {
	ch    byte
	count int
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		ggt := 0
		symbols := make([]Symbol, 0)
		for i := 0; i < len(s); i++ {
			if s[i] == '>' {
				ggt++
			}
			if len(symbols) == 0 {
				symbols = append(symbols, Symbol{s[i], 1})
			} else {
				last := symbols[len(symbols)-1]
				if last.ch == s[i] {
					symbols[len(symbols)-1].count++
					//last.count++
				} else {
					symbols = append(symbols, Symbol{s[i], 1})
				}
			}
		}
		ans := make([]int, 0)
		lst := 1
		gt := ggt
		if s[0] == '>' {
			gt++
		}
		//write(f, symbols, "\n")
		for i := len(symbols) - 1; i >= 0; i-- {
			count := symbols[i].count
			if i == 0 {
				count++
			}

			if symbols[i].ch == '>' {
				for j := 0; j < count; j++ {
					ans = append(ans, lst)
					lst++
				}
			} else {
				for j := 0; j < count; j++ {
					ans = append(ans, gt+count-j)
				}
				gt += symbols[i].count
			}
		}
		for i := len(ans) - 1; i >= 0; i-- {
			write(f, ans[i], " ")
		}
		write(f, "\n")

		gt = ggt
		if s[0] == '>' {
			gt++
		}
		lt := gt + 1
		for i := 0; i < len(s); i++ {
			if i == 0 {
				if s[i] == '<' {
					write(f, lt, lt+1)
					lt += 2
				} else {
					write(f, gt, gt-1)
					gt -= 2
				}
			} else {
				if s[i] == '<' {
					write(f, " ", lt)
					lt++
				} else {
					write(f, " ", gt)
					gt--
				}
			}
		}
		write(f, "\n")
	}
}
