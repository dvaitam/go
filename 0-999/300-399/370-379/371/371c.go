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
	var s string
	var nb, ns, nc, pb, ps, pc, r int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s, &nb, &ns, &nc, &pb, &ps, &pc, &r)
	bc := int64(0)
	sc := int64(0)
	cc := int64(0)
	for i := 0; i < len(s); i++ {
		if s[i] == 'B' {
			bc++
		} else if s[i] == 'S' {
			sc++
		} else {
			cc++
		}
	}
	ans := int64(0)
	cost := bc*pb + sc*ps + cc*pc
	for r > 0 {
		if bc <= nb && sc <= ns && cc <= nc {
			ans++
			nb -= bc
			ns -= sc
			nc -= cc
		} else if nb > 0 || ns > 0 || nc > 0 {
			req := int64(0)
			if nb < bc {
				req += (bc - nb) * pb
			}
			if ns < sc {
				req += (sc - ns) * ps
			}
			if nc < cc {
				req += (cc - nc) * pc
			}
			if req <= r {
				r -= req
				ans++
				used := 0
				if nb < bc {
					nb = 0
				} else {
					if bc > 0 {
						used++
					}
					nb -= bc
				}
				if ns < sc {
					ns = 0
				} else {
					if sc > 0 {
						used++
					}
					ns -= sc
				}
				if nc < cc {
					nc = 0
				} else {
					if cc > 0 {
						used++
					}
					nc -= cc
				}
				if used == 0 {
					break
				}

			} else {
				break
			}
		} else {
			break
		}
	}
	ans += r / cost
	write(f, ans, "\n")
}
