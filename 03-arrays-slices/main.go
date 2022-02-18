package main

import "fmt"

func main() {

	//arrays
	var colors = [3]string{"azul", "verde", "amarelo"}


	fmt.Println("O array possui", len(colors), "posições.")
	fmt.Println("Os values são:", colors)
	fmt.Println("O index 0 é", colors[0], ", o index 1 é", colors[1],"e o index 2 é", colors[2])

	//slices

	var favoriteMusic []string

	favoriteMusic = append(favoriteMusic, "Fiona Apple")
	favoriteMusic = append(favoriteMusic, "Rihanna")
	favoriteMusic = append(favoriteMusic, "Alicia Keys")
	favoriteMusic = append(favoriteMusic, "Tulipa Ruiz")
	
	fmt.Println(favoriteMusic)
	fmt.Println(favoriteMusic[0])

	//to 'delete', append two slices with ranges
	favoriteMusic = append(favoriteMusic[:1], favoriteMusic[2:] ...) //func variadic ... para 1 ou mais valores
	
	fmt.Println("Minhas cantoras:", favoriteMusic)
	//delete()

}