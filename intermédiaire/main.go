package main

import(
	"fmt"
)

func main() {
	slice := make([]int, 10)
	slice[0] = 1
	slice[1] = 2
	slice[3] = 3
	slice[4] = 4
	slice[5] = 5
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
	fmt.Println(slice)

	sliceBis := make([]int, 0)
	sliceBis = append(sliceBis, 0)
	sliceBis = append(sliceBis, 1)
	fmt.Println(sliceBis)

	varMap := make(map[string]string)
	varMap["ma clé"] = "ma valeur"
	fmt.Println(varMap)

	go func(texte string) {
		fmt.Println(texte)
	}("Hello Wolrd")

    var input string
    fmt.Scanln(&input)
    fmt.Println("programme terminé, la goroutine a terminé")

    channel := make(chan int)
    go func(){
    	channel <- 1
	}()
	valueFromChannel := <-channel
	fmt.Println(valueFromChannel)
}