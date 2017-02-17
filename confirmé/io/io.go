package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
)
func main() {
	/* Lecture vec ioutil */
	content, err := ioutil.ReadFile("io.txt")
	if err != nil {
		fmt.Println("erreur à la lecture du fichier", err)
	}
	fmt.Println(string(content))

	/* Lecture avec os et bufio */
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
	file.Close()

	/* Ecriture avec ioutil */

	/* Ecriture avec os */
	newFile, err := os.Create("ioos.txt")
	if err != nil {
		fmt.Println("erreur à la création du fichier", err)
	}
	_, err = newFile.Write([]byte("helloworld"))
	if err != nil {
		fmt.Println("erreur à l'écriture dans le fichier", err)
	}
	newFile.Close()
	/* Ecriture avec bufio */
	newFileBufio, err := os.Create("iobufio.txt")
	if err != nil {
		fmt.Println("erreur à la création du fichier", err)
	}
	writer := bufio.NewWriter(newFileBufio)
    _, err = writer.WriteString("helloworld")
	if err != nil {
		fmt.Println("erreur à l'écriture dans le fichier", err)
	}
	writer.Flush()
	newFileBufio.Close()
}
