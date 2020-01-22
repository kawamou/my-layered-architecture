package main

import (
	"fmt"
	"go-cleanarchitecture-restapi/infrastructure"
)

func main() {
	err := infrastructure.Load()
	if err != nil {
		fmt.Println(err)
	}
	infrastructure.Init()
}