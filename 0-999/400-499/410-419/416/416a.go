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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	greater := false
	gv := 0
	lesser := false
	lv := 0
	for i := 0; i < n; i++ {
		var e, s string
		var v int
		fmt.Fscan(reader, &e, &v, &s)
		if s == "Y" {
			if e == ">" || e == ">=" {
				if e == ">" {
					v++
				}
				if !greater {
					greater = true
					gv = v
				} else {
					gv = max(gv, v)
				}
			} else {
				if e == "<" {
					v--
				}
				if !lesser {
					lesser = true
					lv = v
				} else {
					lv = min(lv, v)
				}
			}
		} else {
			if e == "<" || e == "<=" {
				if e == "<=" {
					v++
				}
				if !greater {
					greater = true
					gv = v
				} else {
					gv = max(gv, v)
				}
			} else {
				if e == ">=" {
					v--
				}
				if !lesser {
					lesser = true
					lv = v
				} else {
					lv = min(lv, v)
				}
			}
		}
	}
	if lesser && greater {
		if lv < gv {
			write(f, "Impossible\n")
		} else {
			write(f, lv, "\n")
		}
	} else if lesser {
		write(f, lv, "\n")
	} else {
		write(f, gv, "\n")
	}
}
