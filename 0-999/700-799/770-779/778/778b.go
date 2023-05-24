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

type Expression struct {
	a, b        int
	val         []byte
	ai, bi, exp string
}

func assign(name int, ex []Expression, j int) {
	if ex[name].exp == "C" {
		return
	}
	if ex[name].val[j] != '-' {
		return
	}
	a := ex[name].a
	b := ex[name].b
	if ex[a].val[j] == '-' {
		assign(a, ex, j)
	}
	if ex[b].val[j] == '-' {
		assign(b, ex, j)
	}
	if ex[name].exp == "AND" {
		if ex[a].val[j] == '1' && ex[b].val[j] == '1' {
			ex[name].val[j] = '1'
		} else {
			ex[name].val[j] = '0'
		}
	} else if ex[name].exp == "OR" {
		if ex[a].val[j] == '0' && ex[b].val[j] == '0' {
			ex[name].val[j] = '0'
		} else {
			ex[name].val[j] = '1'
		}
	} else if ex[name].exp == "XOR" {
		if ex[a].val[j] == '0' && ex[b].val[j] == '0' || ex[a].val[j] == '1' && ex[b].val[j] == '1' {
			ex[name].val[j] = '0'
		} else {
			ex[name].val[j] = '1'
		}
	}

}
func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	var_map := map[string]int{}
	ex := make([]Expression, n+1)
	for i := 0; i < n; i++ {
		var name, eq, val string
		exp := Expression{}
		fmt.Fscan(reader, &name, &eq, &val)
		if val[0] == '0' || val[0] == '1' {
			exp.val = []byte(val)
			exp.exp = "C"
		} else {
			exp.ai = val
			var not, b string
			fmt.Fscan(reader, &not, &b)
			exp.exp = not
			exp.bi = b
			exp.val = make([]byte, m)
		}
		ex[i] = exp
		var_map[name] = i
	}
	var_map["?"] = n
	for i := 0; i < n; i++ {
		if ex[i].exp != "C" {
			ex[i].a = var_map[ex[i].ai]
			ex[i].b = var_map[ex[i].bi]
		}
	}
	ex[n] = Expression{exp: "C", val: make([]byte, m)}
	min_ans := make([]byte, m)
	max_ans := make([]byte, m)
	for i := 0; i < m; i++ {
		ex[n].val[i] = '0'
		for j := 0; j < n; j++ {
			if ex[j].exp != "C" {
				ex[j].val[i] = '-'
			}
		}
		for j := 0; j < n; j++ {
			assign(j, ex, i)
		}
		zero_count := 0
		for j := 0; j < n; j++ {
			if ex[j].val[i] == '1' {
				zero_count++
			}
		}
		ex[n].val[i] = '1'
		for j := 0; j < n; j++ {
			if ex[j].exp != "C" {
				ex[j].val[i] = '-'
			}
		}
		for j := 0; j < n; j++ {
			assign(j, ex, i)
		}
		one_count := 0
		for j := 0; j < n; j++ {
			if ex[j].val[i] == '1' {
				one_count++
			}
		}
		if zero_count < one_count {
			min_ans[i] = '0'
			max_ans[i] = '1'
		} else if one_count < zero_count {
			min_ans[i] = '1'
			max_ans[i] = '0'
		} else {
			min_ans[i] = '0'
			max_ans[i] = '0'
		}
	}
	write(f, string(min_ans), "\n")
	write(f, string(max_ans), "\n")
}
