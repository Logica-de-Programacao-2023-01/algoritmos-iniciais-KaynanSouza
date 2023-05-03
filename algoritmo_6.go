package main

import "fmt"

func main() {
	var num int
	cont := 0
	fmt.Print("digite um numero:")
	fmt.Scan(&num)
	for i := 1; i < num; i++ {
		if num%i == 0 {
			cont++
		}
	}
	if cont == 1 {
		fmt.Print("esse numero eh primo")
	} else {
		fmt.Printf("esse numero neo eh primo, pois ele e divisivel por %d numeros", cont)
	}
}
