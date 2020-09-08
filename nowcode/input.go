package main

import "fmt"

type feature struct {
	x int
	y int
}

func main() {
	n := 0
	dir := make(map[feature]int)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		m := 0
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			head := 0
			fmt.Scan(&head)
			if head == 0 {
				continue
			}
			for k := 0; k < head; k++ {
				x, y := 0, 0
				fmt.Scan(&x, &y)
				count, ok := dir[feature{x, y}]
				if ok {
					dir[feature{x, y}] = count + 1
				} else {
					dir[feature{x, y}] = 1
				}
			}
		}
		max := 0
		for _, val := range dir {
			if val > max {
				max = val
			}
		}
		fmt.Println(max)
	}
}
