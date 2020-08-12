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
	fmt.Println(cantando)

}
