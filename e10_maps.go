package main

import (
	"fmt"
)

func main() {

	var m map[string]int
	fmt.Println(m)

	m = make(map[string]int)
	fmt.Println(m)

	m["key"] = 42
	fmt.Println("buscando por key ", m["key"])

	fmt.Println("mapa con nuevo item ", m)

	delete(m, "key")
	fmt.Println("mapa vacio ", m)

	elem, ok := m["key"] // evalúa si key está presente y la retorna
	fmt.Println("buscando elemento key", elem, ok)

	// map literal
	var e = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println("mapa llave, struc", e)
}

type Vertex struct {
	x, y float64
}
