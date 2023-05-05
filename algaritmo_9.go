package main

import (
	"fmt"
	"math"
)

func main() {
	var c, i, t, n, m, jc float64
	fmt.Print("digite o numero do capital:")
	fmt.Scan(&c)
	fmt.Print("digite o numero do taxa:")
	fmt.Scan(&t)
	fmt.Print("digite o numero do tempo:")
	fmt.Scan(&n)
	jc = c * math.Pow((1+i), n)
	m = c + jc
	fmt.Println("juros compostos:", jc)
	fmt.Print("montante:", m)
}
