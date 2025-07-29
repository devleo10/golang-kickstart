package main

import "fmt"

func main() {
	var username string = "Devleo"
	fmt.Println("Hello, " + username + "!")
	fmt.Printf("Variable is of type %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type:%T\n", isLoggedIn)

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type:%T\n", anotherVariable)
}
