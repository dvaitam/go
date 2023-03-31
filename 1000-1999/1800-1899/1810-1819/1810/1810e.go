// 00
package main

import (
	"bufio"
	"container/heap"
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

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	//fmt.Println("before append", h)
	*h = append(*h, x.(int))
	//fmt.Println("after append", h)
}

func (h *IntHeap) Pop() any {
	old := *h
	//fmt.Println("before pop", h)
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([]int, n+1)
		possibles := make([]int, 0)
		for i := 1; i <= n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] == 0 {
				possibles = append(possibles, i)
			}
		}
		adj := map[int][]int{}
		for i := 0; i < n; i++ {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			if adj[u] == nil {
				adj[u] = make([]int, 0)
			}
			if adj[v] == nil {
				adj[v] = make([]int, 0)
			}
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
		//ok := false
		for i := 0; i < len(possibles); i++ {
			h := &IntHeap{}
			heap.Init(h)
			health := 0
			heap.Push(h, possibles[i])
			for h.Len() > 0 {
				if health < a[(*h)[0]] {
					break
				} else {
					heap.Pop(h)
					health++
				}
			}
		}

	}
}
