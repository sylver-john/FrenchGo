package main

import (
  "./router"
)

func main() {
  router := router.NewRouter()
  router.Run()
}