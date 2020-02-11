package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var people model.App
	people.ID = 100
	fmt.Println(people)
}
