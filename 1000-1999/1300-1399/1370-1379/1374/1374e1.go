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

var n, k int

func assign(a []int64, i int, val int64) {
	if a[i] > 0 {
		a[i] = min(a[i], val)
	} else {
		a[i] = val
	}
}
func index(x int, y int) int {
	return x*n + y
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
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	b := make([]int, n)
	t := make([]int, n)
	ah := &IntHeap{}
	heap.Init(ah)
	bh := &IntHeap{}
	heap.Init(bh)
	abh := &IntHeap{}
	heap.Init(abh)
	ans := int64(0)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i], &a[i], &b[i])
		if a[i] == 1 && b[i] == 1 {
			heap.Push(abh, t[i])

		} else if a[i] == 1 {
			heap.Push(ah, t[i])

		} else if b[i] == 1 {
			heap.Push(bh, t[i])

		}
	}
	for k > 0 {
		if ah.Len() > 0 && bh.Len() > 0 {
			if abh.Len() > 0 {
				if (*abh)[0] >= (*ah)[0]+(*bh)[0] {
					ans += int64(heap.Pop(ah).(int) + heap.Pop(bh).(int))
				} else {
					ans += int64(heap.Pop(abh).(int))
				}
				k--
			} else {
				ans += int64(heap.Pop(ah).(int) + heap.Pop(bh).(int))
				k--
			}
		} else if abh.Len() > 0 {
			ans += int64(heap.Pop(abh).(int))
			k--
		} else {
			break
		}
	}
	if k == 0 {
		write(f, ans, "\n")
	} else {
		write(f, "-1\n")
	}

}
