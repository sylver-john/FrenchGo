# Écrire et lire un fichier

Il existe de nombreuses méthodes pour lire et écrire un fichier.À utiliser selon la taille du fichier, la taille du buffer, le type du fichier, etc...

## Avec le package io/oiutil

Pour lire un fichier avec le package os, il suffit d'utiliser la bonne méthode, et de convertir le contenu selon ce qu'on souhaite car on récupère des octets :
```go
content, err := ioutil.ReadFile("io.txt")
if err != nil {
		fmt.Println("erreur à la lecture du fichier", err)
}
fmt.Println(string(content))
```
Ici je lis un fichier txt qui contient "helloworld", récupère les octets que je cast en string pour l'afficher.

## Avec le package os et bufio

Pour lire un fichier avec le package os, j'ai besoin du package bufio pour lire le contenu :
```go
file, err := os.Open("io.txt")
if err != nil {
	fmt.Println("erreur à l'ouverture du fichier", err)
}
fileStats, err := file.Stat()
if err != nil {
	fmt.Println("erreur à la récupération des informations sur le fichier", err)
}
reader := bufio.NewReader(file)
fileContentBytes, err := reader.Peek(int(fileStats.Size()))
if err != nil {
	fmt.Println("erreur à la lecture du contenu du fichier", err)
}
fmt.Println(string(fileContentBytes))
```
Ici, j'ouvre le fichier avec ``os.Open`` puis je récupère sa longueur pour pouvoir utiliser le Reader fournit par le package bufio
Et je n'oublie pas de fermer le fichier :
```go
file.Close()
```
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
