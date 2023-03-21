package main

import "fmt"

func main() {
	var age string
	fmt.Print("Qual é a sua idade? ")
	fmt.Scan(&age)
	fmt.Println("sua idade", age)
}
