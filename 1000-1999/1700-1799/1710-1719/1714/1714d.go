// 1714D. Color with Occurrences
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	w, p, l int
}

var ans []Pair

func substring(st string, start int, sub string) bool {
	for i := 0; i < len(sub); i++ {
		if start+i >= len(st) {
			return false
		}
		if sub[i] != st[start+i] {
			return false
		}
	}
	return true
}
func backtrack(st string, start int, pairs []Pair, so_far []Pair) {
	if start == len(st) {
		if ans == nil || len(so_far) < len(ans) {
			ans = make([]Pair, len(so_far))
			for i := 0; i < len(so_far); i++ {
				ans[i] = so_far[i]
			}
		}
	}
	greatest := Pair{w: -1, l: -1, p: -1}
	for i := 0; i < len(pairs); i++ {
		if pairs[i].p <= start && pairs[i].p+pairs[i].l > start {
			if greatest.w == -1 || greatest.p+greatest.l < pairs[i].p+pairs[i].l {
				greatest = pairs[i]
			}
		}
	}
	if greatest.w != -1 {
		so_far = append(so_far, greatest)
		backtrack(st, greatest.p+greatest.l, pairs, so_far)
		so_far = so_far[:len(so_far)-1]
	}

}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var st string
		fmt.Fscan(reader, &st)
		var n int
		fmt.Fscan(reader, &n)
		s := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &s[i])
		}
		all_pairs := make([]Pair, 0)
		for i := 0; i < len(st); i++ {
			for j := 0; j < n; j++ {
				if substring(st, i, s[j]) {
					all_pairs = append(all_pairs, Pair{w: j, p: i, l: len(s[j])})
				}
			}
		}
		ans = nil
		so_far := make([]Pair, 0)
		backtrack(st, 0, all_pairs, so_far)
		if ans == nil {
			fmt.Println(-1)
		} else {
			fmt.Println(len(ans))
			for i := 0; i < len(ans); i++ {
				fmt.Println(ans[i].w+1, ans[i].p+1)
			}
		}
	}
}
