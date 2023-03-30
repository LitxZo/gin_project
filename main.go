package main

import (
	"fmt"
	"gin_project/router"
)

func main() {
	router := router.WebStart()
	router.Run()
	fmt.Println("hello")
}
