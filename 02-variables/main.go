package main

import (
	"fmt"
)

func main() {

	var name string = "Taís Araújo"
	const idade = 27
	var firstName, lastName = "Taís", "Araujo"
	hobby := "filmes e séries"
	var favoriteMusic string = "Periclao"
	var num float32 = 15.2

	//fmt.Println(name)
	//fmt.Println(idade)
	//fmt.Println(firstName + " " + lastName)

	fmt.Println("Olá, eu sou ", name, "gosto de ", hobby, "e meu cantor favorito é o ", favoriteMusic)

	fmt.Printf("Meu nome completo é:\n %s %s. E eu tenho %v anos", firstName, lastName, idade)

	fmt.Printf("\n Peso do pássaro é: %v gramas", num)
}
