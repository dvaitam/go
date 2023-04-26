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

type Point struct {
	i, j int
}

func has_won(s []string) (bool, bool) {
	first_won, second_won := false, false
	for i := 0; i < 3; i++ {
		poss := true
		for j := 1; j < 3; j++ {
			if s[i][j] != s[i][j-1] {
				poss = false
				break
			}
		}
		if poss {
			if s[i][0] == 'X' {
				first_won = true
			} else if s[i][0] == '0' {
				second_won = true
			}
		}
	}
	for j := 0; j < 3; j++ {
		poss := true
		for i := 1; i < 3; i++ {
			if s[i][j] != s[i-1][j] {
				poss = false
				break
			}
		}
		if poss {
			if s[0][j] == 'X' {
				first_won = true
			} else if s[0][j] == '0' {
				second_won = true
			}
		}
	}
	poss := true
	for i := 1; i < 3; i++ {
		if s[i][i] != s[i-1][i-1] {
			poss = false
		}
	}
	if poss {
		if s[0][0] == 'X' {
			first_won = true
		} else if s[0][0] == '0' {
			second_won = true
		}
	}
	poss = true
	for i := 1; i < 3; i++ {
		if s[i][2-i] != s[i-1][3-i] {
			poss = false
		}
	}
	if poss {
		if s[0][2] == 'X' {
			first_won = true
		} else if s[0][2] == '0' {
			second_won = true
		}
	}
	return first_won, second_won
}
func main() {
	s := make([]string, 3)
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &s[i])
	}
	crosses := make([]Point, 0)
	noughts := make([]Point, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s[i][j] == 'X' {
				crosses = append(crosses, Point{i: i, j: j})
			} else if s[i][j] == '0' {
				noughts = append(noughts, Point{i: i, j: j})
			}
		}
	}
	c := len(crosses)
	n := len(noughts)
	if c-n > 1 || n > c {
		write(f, "illegal\n")
	} else {
		first_won, second_won := has_won(s)
		if first_won && second_won {
			write(f, "illegal\n")
		} else if first_won {
			if c == n {
				write(f, "illegal\n")
			} else {
				write(f, "the first player won\n")
			}
		} else if second_won {
			if c == n {
				write(f, "the second player won\n")
			} else {
				write(f, "illegal\n")
			}
		} else {
			if c == n {
				write(f, "first\n")
			} else if c+n != 9 {
				write(f, "second\n")
			} else {
				write(f, "draw\n")
			}
		}

	}

}
