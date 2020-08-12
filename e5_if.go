package main

import "fmt"

func main() {
	fmt.Println("basic If ", basicIf(9))
	fmt.Println("asignation and If ", asignationAndIf(42, 2))
}

func basicIf(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func asignationAndIf(b, c int) int {
	if a := b + c; a < 42 {
		return a
	} else {
		return a - 42
	}
}
