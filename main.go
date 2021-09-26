package main

import (
	"log"
)

const port string = ":8080"

func main() {
	log.Println(NewRouter().Run(port))
}
