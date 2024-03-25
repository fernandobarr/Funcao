package main

import "fmt"

func main() {
	fmt.Println("Teste")
	x := 0.0
	numero := func() float64 {

		x++
		return x
	}

	fmt.Println(numero())
	fmt.Println(numero())
}
