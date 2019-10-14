package main

import (
	"fmt"
)

func main() {
	const i = 1000
	for i := 0; i < 12; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Hello World!")
	}

	foo := "100"

	if foo == "100" {

		fmt.Println("Hello World!")

		arr := [5]int{1, 2, 3, 4, 5}

		j := 0
		for j < len(arr) {

			fmt.Println(arr[j])
			j++

		}
		s := make([]string, 3)
		fmt.Println("emp:", s)

		s[0] = "a"
		s[1] = "b"
		s[2] = "c"

		s = append(s, `d`)
		fmt.Println("emp:", s)

		l := s[2:5]
		fmt.Println("sl1:", l)
	}
}
