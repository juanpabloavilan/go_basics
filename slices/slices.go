package slices

import (
	"fmt"
)

func SlicesLab() {

	s := make([]int, 2)
	s[0] = 2

	fmt.Println(s)

	s = append(s, 3, 4)

	fmt.Println(s)

	c := make([]int, len(s))
	copy(c, s)
}
