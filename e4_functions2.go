package main

import "fmt"

func main() {
	fmt.Println(adder(1, 2, 3)) // 6
	fmt.Println(adder(9, 9))    // 18
	nums := []int{10, 20, 30}
	fmt.Println(adder(nums...)) // 60
}

// Usando ... antes del tipo se puede indicar que se aceptan 0 o más parámetros del mismo tipo
func adder(args ...int) int {
	total := 0
	for _, v := range args {
		// Itera según tantos argumentos reciba
		total += v
	}
	return total
}
