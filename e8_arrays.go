package main

import "fmt"

func main() {

	var a [10]int // array tamaño 10 tipo enteros
	fmt.Println(a)

	a[3] = 42 // define elemento
	fmt.Println(a)

	i := a[3] // leer el elemento
	fmt.Println(i)

	// declaración e inicialización
	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	c := [6]int{1, 2, 3, 4, 5, 6} //tamaño definido
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // elipsis -> Compilador define el tamaño
	fmt.Println(d)
}
