package main

import "fmt"

//forma objeto
type videoGame struct {
	name string
	brand string
	year uint
}


//toda instancia da struct vai ter func printInfo
func(v videoGame) printInfo() {
	fmt.Println("Nome:", v.name, "Marca:", v.brand, "Ano:", v.year)
}

func NewVideoGame(name string, brand string) videoGame {
	return videoGame{
		name: name,
		brand: brand,
		year: 1999999,
	}
}

func main() {

	playstation := videoGame{
		name: "Play One",
		brand: "Sony",
		year: 2000,
	}

	superNintendo := videoGame{
		name: "Super Nintendo",
		brand: "Nintendo",
	}

	superNintendo.year = 1990

	fmt.Println(playstation)
	fmt.Println(playstation.year)

	fmt.Println(superNintendo)

	superNintendo.printInfo()


}