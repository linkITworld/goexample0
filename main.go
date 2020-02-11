package main

import (
	"fmt"

	"github.com/linkITworld/goexample0/model"
)

func main() {
	fmt.Println("Hello")
	var people model.App
	people.ID = 100
	people.Name = "strr"
	people.UUID = "sty"
	fmt.Println(people)
}
