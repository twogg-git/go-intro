package main

import "fmt"

var t1 int //(sin inicializar)

var t2 int = 42           //(con inicialización)
var t3, t4 int = 42, 1302 //(declaración múltiple)
var t5 = 42               //(tipo omitido se infiere)

const constant = "Esto es una constante"

func main() {

	fmt.Println("sin inicializar t1", t1)
	fmt.Println("con inicialización t2", t2)
	fmt.Println("declaración múltiple t3 t4", t3, " ", t4)
	fmt.Println("tipo omitido se infiere t5 ", t5)

	t6 := 42 //(tipo omitido, sólo dentro de funciones)
	fmt.Println("tipo omitido t6 ", t6)
}
