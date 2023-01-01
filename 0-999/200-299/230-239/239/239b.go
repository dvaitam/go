// 00
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
}
type Object struct {
	c byte
	v int
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
func print_list(f *bufio.Writer, l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		val := e.Value.(Object)

		write(f, val.c, " ", val.v, " ")
	}
	write(f, "\n")
}
func main() {
	var n, q int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &q)
	var s string
	fmt.Fscan(reader, &s)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		ll := list.New()
		for j := l - 1; j < r; j++ {
			if s[j] == '<' || s[j] == '>' {
				ll.PushBack(Object{c: s[j], v: 0})
			} else {
				ll.PushBack(Object{c: '-', v: int(s[j] - '0')})
			}
		}
		right := true
		curr := ll.Front()
		m := map[int]int{}
		for true {
			val := curr.Value.(Object)
			// print_list(f, ll)
			if val.c == '-' {
				m[val.v]++
			} else {
				if val.c == '<' {
					right = false
				} else {
					right = true
				}
			}
			if right {
				if curr.Next() == nil {
					break
				} else {
					tmp := curr
					curr = curr.Next()
					tmp_val := tmp.Value.(Object)
					curr_val := curr.Value.(Object)
					if tmp_val.c == '-' {
						if tmp_val.v == 0 {
							ll.Remove(tmp)
						} else {
							tmp_val.v--
							tmp.Value = tmp_val
						}
					} else if tmp_val.c != '-' && curr_val.c != '-' {
						ll.Remove(tmp)
					}
				}
			} else {
				if curr.Prev() == nil {
					break
				} else {
					tmp := curr
					curr = curr.Prev()
					tmp_val := tmp.Value.(Object)
					curr_val := curr.Value.(Object)
					if tmp_val.c == '-' {
						if tmp_val.v == 0 {
							ll.Remove(tmp)
						} else {
							tmp_val.v--
							tmp.Value = tmp_val
						}
					} else if tmp_val.c != '-' && curr_val.c != '-' {
						ll.Remove(tmp)
					}
				}
			}

		}
		for k := 0; k < 10; k++ {
			write(f, m[k])
			if k == 9 {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}

	}
}
