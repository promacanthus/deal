package main

import "fmt"

func main() {
	n := 1
	items := make([]int, n)

	fmt.Println(len(items))
	fmt.Println(cap(items))
	fmt.Println(items)
}
