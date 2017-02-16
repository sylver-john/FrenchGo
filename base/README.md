# Les bases

- [Ecrire un fichier go](##ecrire-un-fichier-go)
- [Déclarer et afficher une variables](##déclarer-et-afficher-une-variables)
- [Créer et utiliser une strucure](##créer-et-utiliser-une-strucure)
- [Utiliser un opérateur](##utiliser-un-opérateur)
- [Conditons et boucles](##conditons-et-boucles)

Le fichier main.go contient le code avec les différents exemples.

## Ecrire un fichier go :

Déclarer le package auquel appartient le fichier :
```go
package leNomDeMonDossier
```

Importer un package (car en Go tout est package, la moindre fonctionnalitée doit être importée) :
```go
import(
	"fmt"
)
```

Déclarer une fonction (ici la fonction main, point d'entrée de l'application) :
```go
func main() {

}
```

## Déclarer et afficher une variables.

La forme simple :
```go
	var toto string
	toto = "hello"
	fmt.Println(toto)
```
  
La forme plus rapide :
```go
  tata := "world"
	fmt.Println(tata)
``` 

Les différents types de bases :
  - bool
  - string
  - int, int8, int16, int32, int64
  - uint, uint8, uint16, uint32, uint64, uintptr
  - byte (unit8)
  - rune (int32)
  - float32, float64
  -complex64, complex128
  
## Créer et utiliser une strucure :

Créer sa structure :
```go
type A struct {
	Hello string
	World string
}
```
Déclarer les attributs avec une majuscule permet de les exportés (attribut public)

Accèder à ses attributs :
```go
	a := A{"hello", "world"}
	fmt.Println(a.Hello, a.World)
 ```
 Attention, on peut accèder aux attributs de cette manière seulement si les attributs sont exportés.
 
 Déclarer une structure plus complexe :
 ```go
 type A struct {
	Hello string
	World string
}

type B struct {
	A A
	From string
}
```

Déclarer une constante :
```go
import(
	"fmt"
)

const constante string = "Constante"

func main() {
```

## Utiliser un opérateur

Faire une somme :
```go
var x, y int
x = 1
y = 2
fmt.Println(x + y)
```

Les opérateurs :
```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
&^          &^=
```

## Conditons et boucles

Utiliser un if else :
```go
if 1 != 2 {
	fmt.Println("1 != 2")
} else if 1 == 3 {
	fmt.Println("on a un problème")
} else {
	fmt.Println("on a définitivement un problème")
}
```

Utiliser un switch :
```go
switchTo := 1
switch switchTo {
case 1:
	fmt.Println(switchTo)
case 2:
	fmt.Println(switchTo)
default:
	fmt.Println(switchTo)
}
```

Utiliser une boucle for :
```go
for compteur := 0; compteur < 5; compteur++ {
	fmt.Println(compteur)
}
```
