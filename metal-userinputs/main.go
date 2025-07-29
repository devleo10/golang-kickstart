package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hey there!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	//comma ok || err ok
	input, err:= reader.ReadString('\n')
	fmt.Println("Thanks for rating our Pizza:", input)
    fmt.Println("Error:",err)
}
