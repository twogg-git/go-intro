# Golang ejemplos basicos

## Declaración & Variables
```sh
var foo int (sin inicializar)   
var foo int = 42 (con inicialización)   
var foo, bar int = 42, 1302 (declaración múltiple)   
var foo = 42 (tipo omitido se infiere)   
foo := 42 (tipo omitido, sólo dentro de funciones)   
```
```sh
const constant = "Esto es una constante"   
```
```sh
bool, string, int   
int8 int16 int32 int64   
uint uint8 uint16 uint32 uint64 uintptr   
float32 float64 complex64 complex128   
byte (alias for uint8)   
```

## Funciones

```sh
// función simple
func functionName() {}

// función con parámetros (nombre tipo)
func functionName(param1 string, param2 int) {} 

// múltiples parámetros mismo tipo
func functionName(param1, param2 int) {}

// tipo del retorno inferido
func functionName() int { 
  return 42
}
```

```sh
// retorno con múltiples valores
func returnMulti() (int, string) { 
  return 42, "foobar"
}
var x, str = returnMulti()
```

```sh
// retorno múltiple inferido
func returnMulti2() (n int, s string) { 
  n = 42
  s = "foobar"
  // retornará n y s 
  return
}
var x, str = returnMulti2()
```

```sh
 
func main() { 
  fmt.Println(adder(1, 2, 3)) // 6 
  fmt.Println(adder(9, 9)) // 18 
  nums := []int{10, 20, 30} 
  fmt.Println(adder(nums...)) // 60
}

// Usando ... antes del tipo se puede indicar que se aceptan 0 o más parámetros del mismo tipo
func adder(args ...int) int { 
  total := 0
  for _, v := range args { // Itera según tantos argumentos reciba 
    total += v
  }
  return total 
}
```

##  Estructuras: If
```sh
func main() {
  // Basico
  if x > 0 { 
    return x
  } else {
    return -x
  }

  if a := b + c; a < 42 { 
    return a
  } else {
    return a - 42
  } 
}
```

## Estructuras: Structs
```sh
strings := []string{"hello", "world"}
for i, s := range strings {
	fmt.Println(i, " ", s)
}
````

##  Estructuras: Switch
```sh
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

  // también se pueden hacer comparaciones
  number := 42 
  switch {
  case number < 42: 
    fmt.Println("Smaller")
  case number == 42: 
    fmt.Println("Equal")
  case number > 42: 
    fmt.Println("Greater")
  }
```

## Estructuras: Arrays
```sh
var a [10]int // array tamaño 10 tipo enteros 
a[3] = 42 // define elemento
i := a[3] // leer el elemento
// declaración e inicialización
var a = [2]int{1, 2}
a := [2]int{1, 2} //tamaño definido
a := [...]int{1, 2} // elipsis -> Compilador define el tamaño
```

## Estructuras: Slice
```sh
 var a []int // declaración similar al array, tamaño no especificado 
 var a = []int {1, 2, 3, 4} // declaración con inicialización
a := []int{1, 2, 3, 4} // versión corta
chars := []string{0:"a", 2:"c", 1: "b"} // ["a", "b", "c"]
// creación con función make
a = make([]byte, 5, 5) // primero tamaño, luego capacidad 
a = make([]byte, 5) // capacidad es opcional
// creación con base en array
x := [3]string{"Лайка", "Белка", "Стрелка"} 
s := x[:] // a slice referencing the storage of x
```

## Estructuras: Maps
```sh
var m map[string]int
m = make(map[string]int) m["key"] = 42 fmt.Println(m["key"])
delete(m, "key")
elem, ok := m["key"] // evalúa si key está presente y la retorna

// map literal
var m = map[string]Vertex{
  "Bell Labs": {40.68433, -74.39967}, 
  "Google": {37.42202, -122.08408},
}
```
## Estructuras: Structs
```sh
package main
import "fmt"

type Employee struct { 
  firstName, lastName string
  age, salary int 
}

func main() {
  //creación de estructuras usando el nombre 
  emp1 := Employee{
  firstName: "Sam",
  lastName: "Anderson",
  age: 25,
  salary: 500, 
}
// creación de estructuras sin usar nombres
emp2 := Employee{"Thomas", "Paul", 29, 800} 
fmt.Println("Employee 1", emp1) 
fmt.Println("Employee 2", emp2)
}
```