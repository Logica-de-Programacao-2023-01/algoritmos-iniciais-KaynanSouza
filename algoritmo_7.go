package main

import "fmt"

/*Algoritmo de cálculo de área de um triângulo: um algoritmo que calcula a área de um triângulo com base na base e altura fornecidas.*/

func main() {
	var h, b float64
	fmt.Print("digite a altura:")
	fmt.Scan(&h)
	fmt.Println("digite a base:")
	fmt.Scan(&b)
	fmt.Print("a area do triangulo eh:", (b*h)/2)
}
