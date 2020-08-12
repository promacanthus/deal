package string

import "fmt"

func reverseString(s []byte) {
	n := len(s)
	if n < 1 {
		fmt.Printf("%c\n", s)
		return
	}

	i, j := 0, n-1
	for i < j {
		s[i], s[n-i-1] = s[n-i-1], s[i]
		i++
		j--
	}

	fmt.Print("[")
	for i := 0; i < len(s); i++ {
		fmt.Print("\"")
		fmt.Printf("%c", s[i])
		if i != len(s)-1 {
			fmt.Print("\",")
		} else {
			fmt.Print("\"")
		}

	}
	fmt.Print("]")
}
