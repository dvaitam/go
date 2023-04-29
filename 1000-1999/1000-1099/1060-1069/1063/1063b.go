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

type Node struct {
	i, j, l, r int
}

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	var r, c int
	fmt.Fscan(reader, &r, &c)
	var x, y int
	fmt.Fscan(reader, &x, &y)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	ans := 1
	visited[r-1][c-1] = true
	l := list.New()
	l.PushBack(Node{r - 1, c - 1, 0, 0})
	sl := list.New()
	front := Node{}
	for l.Len() > 0 || sl.Len() > 0 {
		if l.Len() > 0 {
			front = l.Front().Value.(Node)
			l.Remove(l.Front())
		} else {
			front = sl.Front().Value.(Node)
			sl.Remove(sl.Front())
		}
		if front.i > 0 && !visited[front.i-1][front.j] && s[front.i-1][front.j] != '*' {
			ans++
			visited[front.i-1][front.j] = true
			l.PushBack(Node{front.i - 1, front.j, front.l, front.r})
		}
		if front.i+1 < n && !visited[front.i+1][front.j] && s[front.i+1][front.j] != '*' {
			ans++
			visited[front.i+1][front.j] = true
			l.PushBack(Node{front.i + 1, front.j, front.l, front.r})
		}
		if front.j > 0 && !visited[front.i][front.j-1] && s[front.i][front.j-1] != '*' && front.l < x {
			ans++
			visited[front.i][front.j-1] = true
			sl.PushBack(Node{front.i, front.j - 1, front.l + 1, front.r})
		}
		if front.j+1 < m && !visited[front.i][front.j+1] && s[front.i][front.j+1] != '*' && front.r < y {
			ans++
			visited[front.i][front.j+1] = true
			sl.PushBack(Node{front.i, front.j + 1, front.l, front.r + 1})
		}

	}
	write(f, ans, "\n")
}
