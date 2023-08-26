package main

import (
	"fmt"

	"gopl.io/ch2/tempconv"
)

func main() {
	kelvin := tempconv.Celsius(-273.17)
	f := tempconv.CToF(kelvin)
	fmt.Println(f)
}
