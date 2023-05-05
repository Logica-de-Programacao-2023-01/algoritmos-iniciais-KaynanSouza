package main

import "fmt"

func main() {
	var x int
	var produto, media float64
	fmt.Print("digite o tamanho da sua lista:")
	fmt.Scan(&x)
	nota := make([]float64, x)
	peso := make([]float64, x)
	for i := 0; i < len(nota); i++ {
		fmt.Printf("digite a nota %d:", i+1)
		fmt.Scan(&nota[i])
		fmt.Printf("digite o peso %d:", i+1)
		fmt.Scan(&peso[i])
		produto += nota[i] * peso[i]
	}
	media = produto / float64(x)
	fmt.Print(media)
}
