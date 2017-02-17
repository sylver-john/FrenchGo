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

[GIN-debug] GET    /helloworld               --> _/C_/Users/simhoff/FrenchGo/confirm%c3%a9/http/src/handler.HandlerRequest (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

