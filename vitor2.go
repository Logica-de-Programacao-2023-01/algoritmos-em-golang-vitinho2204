package main

import "fmt"

func main() {
	var peso string
	fmt.Print("Qual é o seu peso? ")
	fmt.Scan(&peso)
	fmt.Println("voce tem", peso)
}
