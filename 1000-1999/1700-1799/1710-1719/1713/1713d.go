// 00
package main

import (
	"bufio"
	"fmt"
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
	//reader := bufio.NewReader(os.Stdin)
	//	f := bufio.NewWriter(os.Stdout)
	//defer f.Flush()
	//fmt.Fscan(reader, &T)
	fmt.Scan(&T)
	for t := 1; t <= T; t++ {
		var n int
		var w int
		//fmt.Fscan(reader, &n)
		fmt.Scan(&n)
		a := make([]int, 1<<n)
		for i := 1; i <= 1<<n; i++ {
			a[i-1] = i
		}
		break_three := false
		for len(a) > 1 {
			if break_three {
				break
			}
			if len(a) == 2 {
				fmt.Println("? ", a[0], a[1])

				//	fmt.Fscan(reader, &w)
				fmt.Scan(&w)
				if w == 1 {
					a = []int{a[0]}
				} else if w == 2 {
					a = []int{a[1]}
				} else {
					break_three = true
					break

				}
			} else {
				nxt := make([]int, 0)
				for i := 0; i < len(a); i += 4 {
					fmt.Println("? ", a[i], a[i+3])
					//	fmt.Fscan(reader, &w)
					fmt.Scan(&w)

					if w == 0 {
						fmt.Println("? ", a[i+1], a[i+2])
						//	fmt.Fscan(reader, &w)
						fmt.Scan(&w)

						if w == 1 {
							nxt = append(nxt, a[i+1])
						} else if w == 2 {
							nxt = append(nxt, a[i+2])
						} else {
							break_three = true
							break
						}
					} else if w == 1 {
						fmt.Println("? ", a[i], a[i+2])
						//		fmt.Fscan(reader, &w)
						fmt.Scan(&w)

						if w == 1 {
							nxt = append(nxt, a[i])
						} else if w == 2 {
							nxt = append(nxt, a[i+2])
						} else {
							break_three = true
							break
						}
					} else if w == 2 {
						fmt.Println("? ", a[i+3], a[i+1])
						//fmt.Fscan(reader, &w)
						fmt.Scan(&w)

						if w == 1 {
							nxt = append(nxt, a[i+3])
						} else if w == 2 {
							nxt = append(nxt, a[i+1])
						} else {
							break_three = true
							break
						}
					} else {
						break_three = true
						break
					}
				}
				a = nxt
			}
		}
		if len(a) > 0 {
			fmt.Println("! ", a[0])

		}
	}
}
