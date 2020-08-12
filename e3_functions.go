package main

import "fmt"

func main() {

	fmt.Println("functionInferido", functionInferido())

	var x, str = returnMulti()
	fmt.Println("returnMulti x ", x, "str ", str)

	var x2, str2 = returnMulti2()
	fmt.Println("returnMulti2 x2 ", x2, "str2 ", str2)
}

// función simple
func functionSimple() {
}

// función con parámetros (nombre tipo)
func functionParametros(param1 string, param2 int) {

}

// múltiples parámetros mismo tipo
func functionParametros2(param1, param2 int) {

}

// tipo del retorno inferido
func functionInferido() int {
	return 42
}

// retorno con múltiples valores
func returnMulti() (int, string) {
	return 42, "foobar"
}

// retorno múltiple inferido
func returnMulti2() (n int, s string) {
	n = 42
	s = "foobar"
	return // retornará n y s
}
