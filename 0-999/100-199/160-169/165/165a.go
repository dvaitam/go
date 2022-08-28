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

type Point struct {
	x, y, index int
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	points := make([]Point, n)
	neigbhours := make([]int, n)
	mx := make(map[int][]Point)
	my := make(map[int][]Point)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		points[i] = Point{x: x, y: y, index: i}
		mx[x] = append(mx[x], points[i])
		my[y] = append(my[y], points[i])
	}
	// fmt.Println(mx, my)
	for _, val := range mx {
		if len(val) > 1 {
			//	fmt.Println("processing", val)
			min := val[0].y
			max := val[0].y
			for i := 0; i < len(val); i++ {
				if val[i].y < min {
					min = val[i].y
				}
				if val[i].y > max {
					max = val[i].y
				}
			}
			//	fmt.Println(min, max)
			for i := 0; i < len(val); i++ {
				if val[i].y != min {
					neigbhours[val[i].index]++
				}
				if val[i].y != max {
					neigbhours[val[i].index]++
				}
			}
		}
	}
	for _, val := range my {
		if len(val) > 1 {
			min := val[0].x
			max := val[0].x
			for i := 0; i < len(val); i++ {
				if val[i].x < min {
					min = val[i].x
				}
				if val[i].x > max {
					max = val[i].x
				}
			}
			for i := 0; i < len(val); i++ {
				if val[i].x != min {
					neigbhours[val[i].index]++
				}
				if val[i].x != max {
					neigbhours[val[i].index]++
				}
			}
		}
	}
	// fmt.Println(neigbhours)
	ans := 0
	for i := 0; i < n; i++ {
		if neigbhours[i] == 4 {
			ans++
		}
	}
	fmt.Println(ans)
}
