# Connaissances intermédiaires

Le fichier main.go contient le code avec les différents exemples.

## Les slices

Déclarer une slice :
```go
slice := make([]int, 10)
```
la taille en deuxième paramètre n'est pas obligatoire, une slice peut grandir.

Ajouter un élément dans une slice :
```go
slice[0] = 1
slice[1] = 2
slice[3] = 3
slice[4] = 4
slice[5] = 5
```
ou
```go
slice = append(slice, 6)
slice = append(slice, 7)
slice = append(slice, 8)
slice = append(slice, 9)
slice = append(slice, 10)
```

## Les maps 

Déclarer et utiliser une map:
```go
varMap := make(map[string]string)
varMap["ma clé"] = "ma valeur"
fmt.Println(varMap)
```
## Les goroutines

Une goroutine est un thread d'exécution ultra léger qui permet de mettre en place de la programmation concurrentielle.
Déclarer une goroutine anonyme:
```go
go func(){
}()
```
On peut aussi lancer une fonction existante dans une goroutine avec le mot clé ``go``.
Pour l'utiliser par exemple :
```go
func main() {
	go func(texte string) {
		fmt.Println(texte)
	}("Hello Wolrd")
}
```
L'exemple ci-dessus ne va rien afficher car le treah d'exécution principal va continuer et le programme va se terminer. Pour y remèdier on peut placer un blocage artificiel par exemple attendre que l'on presse une touche du clavier :
```go
func main() {
	go func(texte string) {
		fmt.Println(texte)
	}("Hello Wolrd")
  var input string
  fmt.Scanln(&input)
  fmt.Println("programme terminé, la goroutine a terminé")
}
```

## Les channels

Une channel permet de communiquer entre goroutines.
Déclarer une channel :
```go
channel := make(chan int)
 ```
Dans l'exemple suivant, on nourrit la channel avec un entier à partir d'un second thread d'exécution, que l'on consomme ensuite dans le thread principal :
 ```go
channel := make(chan int)
go func(){
	channel <- 1
}()
valueFromChannel := <-channel
fmt.Println(valueFromChannel)
```

## Les interfaces 

Déclarer une interface :
```go
type C interface {
	String() string
}
```
Implémenter une interface :
```go
type A struct {
	Hello string
}

type B struct {
	World string
}

func (a A) String() string {
	return a.Hello
}

func (b B) String() string {
	return b.World
}

func HelloWorld(hw C) {
	fmt.Println(hw.String())
}
```
Pour implémenter une interface il suffit d'implémenter les méthodes de l'interface aux sutructures. Puis on peut ensuite utiliser les méthodes de l'interface :
```go
hello := A{"hello"}
world := B{"world"}
HelloWorld(hello)
HelloWorld(world)
```

## Les asserts types

L'interface est un objet générique, de nombreuses méthodes retournes des interfaces, pour pouvoir les utiliser il faut les convertir avec l'assert type (ce qui est différent de caster). Par exemple :
```go
var inter interface{}
inter = "hello"
fmt.Println(inter + "world")
```
va lever une erreur, il faut utiliser l'assert type de la valeur contenu dans l'interface :
```go
var inter interface{}
inter = "hello"
fmt.Println(inter.(string) + "world")
```
