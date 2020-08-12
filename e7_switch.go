package main

import "fmt"

func main() {
	switch1("darwin")
	switch2()
}

func switch1(operatingSystem string) {
	fmt.Println("switch1")
	switch operatingSystem {
	case "darwin":
		fmt.Println("Mac OS Hipster")
		// cierra el case automáticamente
	case "linux":
		fmt.Println("Linux Geek")
	default:
		// Windows, BSD, ...
		fmt.Println("Other")
	}
}

func switch2() {
	fmt.Println("switch2")
	number := 42
	switch {
	case number < 42:
		fmt.Println("Smaller")
	case number == 42:
		fmt.Println("Equal")
	case number > 42:
		fmt.Println("Greater")
	}
}
