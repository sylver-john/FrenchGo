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

