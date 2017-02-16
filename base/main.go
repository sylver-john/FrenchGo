package main

import(
	"fmt"
)

const constante string = "Constante"

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

	var x, y int
	x = 1
	y = 2
	fmt.Println(x + y)

	/*
		Les opérateurs de bases :
		+    &     +=    &=     &&    ==    !=    (    )
		-    |     -=    |=     ||    <     <=    [    ]
		*    ^     *=    ^=     <-    >     >=    {    }
		/    <<    /=    <<=    ++    =     :=    ,    ;
		%    >>    %=    >>=    --    !     ...   .    :
		&^          &^=
	*/

	fmt.Println(constante)

    for compteur := 0; compteur < 5; compteur++ {
        fmt.Println(compteur)
    }

    if 1 != 2 {
    	fmt.Println("1 != 2")
    } else if 1 == 3 {
    	fmt.Println("on a un problème")
    } else {
    	fmt.Println("on a définitivement un problème")
    }

    switchTo := 1
    switch switchTo {
    case 1:
    	fmt.Println(switchTo)
    case 2:
    	fmt.Println(switchTo)
	default:
		fmt.Println(switchTo)
   	}


}

type A struct {
	Hello string
	World string
}

type B struct {
	A A
	From string
}