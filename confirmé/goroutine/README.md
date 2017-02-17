# Utiliser les goroutines

## thread pool (voir le fichier pool.go)

Dans cet exemple, j'utilise 10 goroutines pour exécuter un travail qui devrait prendre au minimum 50 secondes.
Tout d'abord on crée notre worker, c'est lui qui exécute la tâche (ça pourrait être un traitement de données, une requête SQL, une lecture de fichier, etc...) :
```go 
func worker(workerId int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", workerId, "commence la tâche :", job)
		time.Sleep(time.Second)
		fmt.Println("worker", workerId, "a terminé la tâche", job)
		results<- job
	}
}
```
Ensuite il faut mettre en place l'exécution de ces workers.
On commence par créer une channel pour y placer les tâches à faire et une pour récupèrer le résultat :
```go
	jobs := make(chan int, 100)
	results := make(chan int, 100)
 ```
Puis on lance nos workers :
```go
for i := 0; i < 10; i++ {
		go worker(i, jobs, results)
}
```
mais ils ne peuvent rien faire pour le moment car la channel contenant les tâches est vide. Il faut les ajouter :
```go
for j := 0; j < 50; j++ {
		jobs<- j
}
close(jobs)
```
Enfin pour lancer la machine, il faut consommer les tâches se trouvant dans la channel des résultats :
```go
for k := 0; k < 50; k++ {
		<-results
}
close(results)
```
Bilan :
On exécute 50 tâches prenant chacunes 1 seconde d'exécution à l'aide de 10 workers, le temps total d'exécution est de moins de 6 secondes, au lieu de 50 secondes si le travail avait été fait de manière séquentielle.

