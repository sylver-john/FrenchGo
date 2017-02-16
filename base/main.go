package main

import(
	"fmt"
)

func main() {
	var toto string
	toto = "hello"
	fmt.Println(toto)

	tata := "world"
	fmt.Println(tata)

	/*
		Les types de bases :
			- bool
			- string
			- int, int8, int16, int32, int64
			- uint, uint8, uint16, uint32, uint64, uintptr
			- byte (unit8)
			- rune (int32)
			- float32, float64
			-complex64, complex128
	*/

	a := A{"hello", "world"}
	fmt.Println(a.Hello, a.World)

	b:= B{a, "from france"}
	fmt.Println(b.A.Hello, b.A.World, b.From)

}

type A struct {
	Hello string
	World string
}

type B struct {
	A A
	From string
}