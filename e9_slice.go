package main

import "fmt"

func main() {
	var a []int // declaración similar al array, tamaño no especificado
	fmt.Println(a)
	var b = []int{1, 2, 3, 4} // declaración con inicialización
	fmt.Println(b)
	c := []int{1, 2, 3, 4} // versión corta
	fmt.Println(c)
	chars := []string{0: "a", 2: "c", 1: "b"} // ["a", "b", "c"]
	fmt.Println(chars)
	// creación con función make
	var d = make([]byte, 5, 5) // primero tamaño, luego capacidad
	fmt.Println(d)
	var e = make([]byte, 5) // capacidad es opcional
	fmt.Println(e)
	// creación con base en array
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	fmt.Println(x)
	s := x[:] // a slice referencing the storage of x
	fmt.Println(s)
}
