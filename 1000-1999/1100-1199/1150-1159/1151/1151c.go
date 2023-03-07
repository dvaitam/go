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
	var l, r int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &l, &r)
	curr := int64(1)
	total := int64(0)
	sm := int64(0)
	odd := true
	odd_numbers := int64(0)
	even_numbers := int64(0)
	for total < r {
		if total+curr < l {
			if odd {
				odd_numbers += curr
			} else {
				even_numbers += curr
			}
			total += curr
			curr = curr * 2
			odd = !odd
			continue
		} else {
			start := max(l, total+1)
			end := min(total+curr, r)
			//write(f, "start ", start, "end ", end, "\n")
			start_number := 2*even_numbers + 2

			if odd {
				start_number = 2*odd_numbers + 1
			}
			range_start := start_number + (start-total-1)*2
			range_end := start_number + (end-total-1)*2
			//write(f, "start number ", start_number, "\n")
			//write(f, "range start ", range_start, " rage end ", range_end, "\n")

			length := ((range_end - range_start) / 2) + 1
			length = length % 1000000007
			sum_numbers := (range_start + range_end) / 2
			sum_numbers = sum_numbers % 1000000007
			sm += length * sum_numbers
			sm = sm % 1000000007
			if odd {
				odd_numbers += curr
			} else {
				even_numbers += curr
			}
			total += curr
			curr = curr * 2
			odd = !odd
		}
	}
	write(f, sm, "\n")
}
