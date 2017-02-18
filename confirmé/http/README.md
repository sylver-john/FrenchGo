# Mettre en place un serveur web

## From scratch 

Go dispose d'un package natif ``net/http`` qui permet de mettre en place assez facilement un serveur web :

```go
package main

import(
	"net/http"
	"log"
)

func main() {
	log.Print("launch the server")
  http.HandleFunc("/", handlerRequest)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handlerRequest(w http.ResponseWriter, r *http.Request) {
	log.Print("=====doing the job======")
	w.Write([]byte{1, 2})
}
```

 Ici on lance  un serveur web sur localhost:1234 qui nous retourne {1, 2} sur la page / .
 Ce package suffit pour créer un petit serveur web et servir des fichiers statics (Go dispose aussi d'un outil de templating) mais devient insuffisant si l'on souhaite mettre en place un plus gros serveur, comme une API (ce qui reste pour moi la meilleur utilisation de Go).
 
 Pour se faire on peut utiliser un framework, il en existe de nombreux :
 - [Gin](https://github.com/gin-gonic/gin)
 - [Martini](https://github.com/go-martini/martini)
 - [Echo](https://github.com/labstack/echo)
 - [Revel](https://github.com/revel/revel)
 - etc...
 
Pour mon exemple j'utilise Gin.
 
## Un exemple d'API avec Gin
 
Tout d'abord j'utilise la strucure suivante :
 ```
 main.go
 	- handler
		- handler.go
	- router
		- router.go
		- routes.go
	- services
		- sqlManager.go
```

Cette structure permet de bien séparer les rôles (MVC) et garantit d'avoir une application maintenable en cas d'évolution.
Pour pouvoir lancer l'application, il faut tout d'abord installer le package Gin :
```go
go get github.com/gin-gonic/gin
```
Puis lancer le serveur :
```go 
go run src/main.go
```
Si le serveur fonctionne vous devez voir les informations Gin :
```go
λ go run src\main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /helloworld               --> _/C_/***/http/src/handler.HandlerRequest (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```
### Le routeur 

Comme on peut le voir dans le fichier ``main.go``, je n'utilise pas directement le router Gin, mais un routeur maison, ça permet de rajouter une couche d'abstraction, et de pouvoir ajouter des fonctionnalités plus facilement d'après moi. 
```go
func main() {
  router := router.NewRouter()
  router.Run()
}

func NewRouter() *gin.Engine {

	router := gin.Default()
	for _, route := range routes {
		var handler gin.HandlerFunc

		handler = route.HandlerFunc
		router.Handle(route.Method, route.Pattern, handler)
	}
	return router
}
```
Ensuite pour pouvoir bien lister les routes que sert notre API, j'utilise le fichier ``routes.go`` qui contient la liste des routes :
```go
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HandlerRequest",
		"GET",
		"/helloworld",
		handler.HandlerRequest,
	},
}
```

### Le handler

Le handler se retrouve aussi séparer du reste dans le fichier ``handler.go`` :
```go
func HandlerRequest(c *gin.Context) {
	services.DoInsert()	
  	c.JSON(http.StatusOK, "hello world!")
}
```
C'est ici que l'on récupère la requête HTTP, que l'on effectue le traitement et que l'on retourne la réponse HTTP. Par exemple un POST pour ensuite faire une insertion en base de données.
