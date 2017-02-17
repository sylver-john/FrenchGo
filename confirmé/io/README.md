# Écrire et lire un fichier

Il existe de nombreuses méthodes pour lire et écrire un fichier.À utiliser selon la taille du fichier, la taille du buffer, le type du fichier, etc...

## Avec le package os

Pour lire un fichier avec le package os, il suffit d'utiliser la bonne méthode, et de convertir le contenu selon ce qu'on souhaite car on récupère des octets :
```go
content, err := ioutil.ReadFile("io.txt")
if err != nil {
		fmt.Println("erreur à la lecture du fichier", err)
}
fmt.Println(string(content))
```
Ici je lis un fichier txt qui contient "helloworld", récupère les octets que je cast en string pour l'afficher.

Pour écrire un fichier, je commence par le créer (si il n'existe pas, sinon on fait une lecture du ficher avant) puis je cast ma chaîne de caracère en octets (chemin inverse de la lecture donc) avant de l'écrire dans le fichier :
```go
newFile, err := os.Create("iobis.txt")
if err != nil {
	fmt.Println("erreur à la création du fichier", err)
}
_, err = newFile.Write([]byte("helloworld")
if err != nil {
	fmt.Println("erreur à l'écriture dans le fichier", err)
}
```
Et je n'oublie pas de fermer le fichier : 
```go
file.Close()
```

## Avec le package io/oiutil et bufio

