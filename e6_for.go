package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("for: tradicional")
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	fmt.Println("\nfor: solo la condición")
	var a int
	for a < 10 { // a := 0; a < 10; a++
		fmt.Print(a)
		a++
	}

	fmt.Println("\nfor: similar a while")
	f := 1
	for f < 10 {
		f *= 2
		fmt.Print(f)
	}

	fmt.Println("\nfor: mientras sea verdadero")
	cantando := ""
	for !strings.Contains(cantando, "pato,pato,pato,") {
		//cantando += "pato,"
		cantando = cantando + "pato,"
		fmt.Println(cantando)
	}
	cantando += "ganzo!!!"
	fmt.Print(cantando)

	fmt.Println("\nfor: iterador de valores")
	strings := []string{"hola", "mundo", "como", "estas?"}
	for i, s := range strings {
		fmt.Println(i, " ", s)
	}

	sum := 0
	for t := 1; t < 5; t++ {
		if t%2 != 0 { // omite numeros impares
			continue // salta esta ejecución

		}
		// break terminaria el for completo
		sum += t
	}
	fmt.Println(sum) // 6 (2+4)

}
