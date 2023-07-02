package main

import (
	"fmt"

	"github.com/basp/luna"
)

func main() {
	u := luna.Point(1, 2, 3)
	v := luna.Point(2, 3, 4)
	w := u.Add(v)
	fmt.Println(w)
}
