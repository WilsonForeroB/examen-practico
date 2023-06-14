package main

import "fmt"

func suma(x, y int) (w int) {
	w = x + y
	return w
}

func resta(x, y int) (z int) {
	z = x - y
	return z
}

func main() {

	var (
		x int
		y int
	)

	fmt.Print("Digite el primer numero entero: ")
	fmt.Scanln(&x)
	fmt.Print("Digite el segundo numero entero: ")
	fmt.Scanln(&y)
	fmt.Println("La suma is:", suma(x, y))
	fmt.Println("La resta is:", resta(x, y))
}
