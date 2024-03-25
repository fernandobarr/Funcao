package main

import "fmt"

func inicial(xPonteiro *int) {
	*xPonteiro = 3
}

func main() {

	x := 5
	inicial(&x)
	fmt.Println(x)
}
