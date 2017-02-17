package main

import(
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < 10; i++ {
		go worker(i, jobs, results) // on lance 10 workers, qui ne font rien car pas de tâche pour le moment
	}

	for j := 0; j < 50; j++ {
		jobs<- j // on nourrit la channel avec 10 tâches à faire
	}
	close(jobs)
	for k := 0; k < 50; k++ {
		<-results // on consomme les résultats de la channel
	}
	close(results)
	fmt.Println(time.Now())
}

func worker(workerId int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", workerId, "commence la tâche :", job)
		time.Sleep(time.Second)
		fmt.Println("worker", workerId, "a terminé la tâche", job)
		results<- job
	}
}