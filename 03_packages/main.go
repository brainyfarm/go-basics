package main

import (
	"fmt"
	"math"

	"github.com/brainyfarm/go-basics/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))
	fmt.Println(math.Remainder(4, 3))
	fmt.Println(strutil.Reverse("olleh"))
}
