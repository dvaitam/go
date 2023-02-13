package main

import "fmt"

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Println("This is Zero.")
	}
	go func() {
		data++
	}()
	if data == 1 {
		fmt.Println("This is one.")
	}
	go func() {
		data++
	}()
	if data == 2 {
		fmt.Println("This is two.")
	}
}
