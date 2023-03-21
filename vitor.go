package main

import "fmt"

func main() {
	var name string
	fmt.Print("Qual é o seu nome? ")
	fmt.Scan(&name)
	fmt.Println("Olá", name)
}
