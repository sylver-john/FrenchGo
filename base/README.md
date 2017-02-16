# Les bases

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

  
